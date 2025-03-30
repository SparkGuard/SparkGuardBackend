package db

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"golang.org/x/crypto/argon2"
	"log"
	"strings"
)

// Параметры Argon2
const (
	Argon2Time      = 3         // Количество итераций
	Argon2Memory    = 64 * 1024 // Потребление памяти (64 MB)
	Argon2Threads   = 4         // Число потоков
	Argon2KeyLength = 32        // Размер хэша (в байтах)
)

// Генерация Argon2-хэша пароля
func HashPassword(password string) (string, error) {
	// Генерация соли
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	// Хэширование с Argon2
	hash := argon2.IDKey([]byte(password), salt, Argon2Time, Argon2Memory, Argon2Threads, Argon2KeyLength)

	// Кодирование соли и хэша в формат Base64
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Храним результат в стандартном формате $argon2id$
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", Argon2Memory, Argon2Time, Argon2Threads, base64Salt, base64Hash), nil
}

func parseHash(hash string) (uint32, uint32, uint8, []byte, []byte, error) {
	var memory, time uint32
	var threads uint8
	var base64Salt, base64Hash string

	// Split the hash using '$' as the separator
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return 0, 0, 0, nil, nil, errors.New("invalid hash format")
	}

	// Extract parameters
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return 0, 0, 0, nil, nil, fmt.Errorf("failed to parse parameters: %v", err)
	}

	base64Salt, base64Hash = parts[4], parts[5]

	// Decode Base64-encoded data
	salt, err := base64.RawStdEncoding.DecodeString(base64Salt)
	if err != nil {
		return 0, 0, 0, nil, nil, fmt.Errorf("failed to decode salt: %v", err)
	}
	hashBytes, err := base64.RawStdEncoding.DecodeString(base64Hash)
	if err != nil {
		return 0, 0, 0, nil, nil, fmt.Errorf("failed to decode hash: %v", err)
	}

	return memory, time, threads, salt, hashBytes, nil
}

// Проверка пароля на основе хранимого Argon2-хэша
func VerifyPassword(password, hash string) (bool, error) {
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash: %s\n", hash)

	memory, time, threads, salt, hashBytes, err := parseHash(hash)
	if err != nil {
		fmt.Printf("Error parsing hash: %v\n", err)
		return false, err
	}

	// Генерация хэша для введенного пароля
	computedHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(hashBytes)))

	fmt.Printf("Expected hash: %x\n", hashBytes)
	fmt.Printf("Computed hash: %x\n", computedHash)

	// Сравнение хэшей
	if subtle.ConstantTimeCompare(hashBytes, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}

func GetUsers() ([]*User, error) {
	var users = make([]*User, 0)

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "email", "access_level", "password").From("users")

	query, args := sb.Build()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccessLevel, &user.Password); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

// Создание нового пользователя
func CreateUser(user *User, password string) error {
	sb := sqlbuilder.PostgreSQL.NewInsertBuilder()

	// Хэширование пароля
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Добавление данных нового пользователя
	sb.InsertInto("users").Cols("name", "email", "access_level", "password").Values(
		user.Name, user.Email, user.AccessLevel, hashedPassword).SQL("RETURNING id")

	query, args := sb.Build()

	return db.QueryRow(query, args...).Scan(&user.ID)
}

// Получение пользователя по ID
func GetUser(id uint) (*User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "email", "access_level", "password").From("users").Where(sb.Equal("id", id))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	user := User{}

	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccessLevel, &user.Password); err != nil {
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}

	return &user, nil
}

// Проверка логина и пароля
func VerifyLogin(email, password string) (*User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id", "name", "email", "access_level", "password").From("users").Where(sb.Equal("email", email))

	query, args := sb.Build()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	user := User{}
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AccessLevel, &user.Password); err != nil {
		return nil, err
	}

	// Проверка пароля
	match, err := VerifyPassword(password, user.Password)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func DeleteUser(id uint) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	sb.DeleteFrom("users").Where(sb.Equal("id", id))
	query, args := sb.Build()
	_, err := db.Exec(query, args...)
	return err
}
