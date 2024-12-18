package repositoriesusers

import (
    "fmt"
    "github.com/karlseguin/ccache"
    "time"
    "users-api/dao"
)

type CacheConfig struct {
    TTL time.Duration // Cache expiration time
}

type Cache struct {
    client *ccache.Cache
    ttl    time.Duration
}

func NewCache(config CacheConfig) Cache {
    // Initialize ccache with default settings
    cache := ccache.New(ccache.Configure())
    return Cache{
        client: cache,
        ttl:    config.TTL,
    }
}

// Helper functions to generate cache keys
func cacheKeyByID(id int64) string {
    return fmt.Sprintf("user:id:%d", id)
}

func cacheKeyByUsername(username string) string {
    return fmt.Sprintf("user:username:%s", username)
}

func (repository Cache) GetAll() ([]users.User, error) {
    return nil, fmt.Errorf("GetAll not implemented in cache")
}

func (repository Cache) GetUserByID(id int64) (users.User, error) {
    idKey := cacheKeyByID(id)
    item := repository.client.Get(idKey)
    if item != nil && !item.Expired() {
        user, ok := item.Value().(users.User)
        if !ok {
            return users.User{}, fmt.Errorf("cached value is not of type User")
        }
        return user, nil
    }
    return users.User{}, fmt.Errorf("cache miss for user ID %d", id)
}

func (repository Cache) GetByUsername(username string) (users.User, error) {
    userKey := cacheKeyByUsername(username)
    item := repository.client.Get(userKey)
    if item != nil && !item.Expired() {
        user, ok := item.Value().(users.User)
        if !ok {
            return users.User{}, fmt.Errorf("cached value is not of type User")
        }
        return user, nil
    }
    return users.User{}, fmt.Errorf("cache miss for username %s", username)
}

func (repository Cache) Create(user users.User) (int64, error) {
    idKey := cacheKeyByID(user.UserID)
    userKey := cacheKeyByUsername(user.Username)
    repository.client.Set(idKey, user, repository.ttl)
    repository.client.Set(userKey, user, repository.ttl)
    return user.UserID, nil
}

func (repository Cache) Update(user users.User) error {
    idKey := cacheKeyByID(user.UserID)
    userKey := cacheKeyByUsername(user.Username)
    repository.client.Set(idKey, user, repository.ttl)
    repository.client.Set(userKey, user, repository.ttl)
    return nil
}

func (repository Cache) Delete(id int64) error {
    idKey := cacheKeyByID(id)
    user, err := repository.GetUserByID(id)
    if err != nil {
        return fmt.Errorf("error retrieving user by ID for deletion: %w", err)
    }
    repository.client.Delete(idKey)
    userKey := cacheKeyByUsername(user.Username)
    repository.client.Delete(userKey)
    return nil
}

func (repository Cache) GetUserByEmail(email string) (users.User, error) {
    return users.User{}, fmt.Errorf("GetUserByEmail not implemented in cache")
}



func (repository Cache) InscriptionCourses(userID int64, courseID string) error {
    // Generar la clave para la inscripción
    key := fmt.Sprintf("inscription:%d:%s", userID, courseID)
    
    // Crear el valor de la inscripción (puede ser cualquier valor, aquí usamos una cadena vacía)
    repository.client.Set(key, "", repository.ttl)

    // Guardar la inscripción en el caché
    return nil
}
