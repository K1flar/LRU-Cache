# LRU Cache

Потокобезопасный менеджер кеша на основе LRU (Least Recently Used) 

## Установка
1. Установка зависимости в вашу Go программу:
```bash
go get -u github.com/K1flar/LRU-Cache
```
2. Подключение в коде:
```go
import "github.com/K1flar/LRU-Cache/pkg/lru"
```

## Алфавитный указатель

1. [func Cap](README.md#func-cap)
2. [func Delete](README.md#func-delete)
3. [func Exist](README.md#func-exist)
4. [func FlushAll](README.md#func-flushall)
5. [func Get](README.md#func-get)
6. [func Keys](README.md#func-keys)
7. [func Len](README.md#func-len)
8. [func LoadJSON](README.md#func-loadjson)
9. [func New](README.md#func-new)
10. [func Rename](README.md#func-rename)
11. [func Resize](README.md#func-resize)
12. [func SaveJSON](README.md#func-savejson)
13. [func Set](README.md#func-set)
14. [func Values](README.md#func-values)

## Функции

### [func Cap](#func-cap)
``` 
func Cap() int
```
Cap возвращает ёмкость кеша.

### [func Delete](#func-delete)
```
func Delete(key K) error
```
Delete удаляет элемент с ключём ``key``.

### [func Exist](#func-exist)
```
func Exist(key K) bool
```
Exist проверяет на наличие элемент с ключём ``key``.

### [func FlushAll](#func-flushall)
```
func FlushAll() error
```
FlushAll очищает кеш.

### [func Get](#func-get)
```
func Get(key K) (value V, ok bool)
```
Get возвращает элемент с ключём ``key`` если он есть.

### [func Keys](#func-keys)
```
func Keys() []K
```
Keys возвращает все ключи в кеше.

### [func Len](#func-len)
```
func Len() int
```
Len возвращает количество элементов в кеше.

### [func LoadJSON](#func-loadjson)
```
func LoadJSON(filePath string) error
```
LoadJSON загружает данные в формате JSON из файла ``filePath`` в кеш.

### [func New](#func-new)
```
func New[K comparable, V any](cap int) *LRUCache
```
New создает кеш размером ``cap``.

### [func Rename](#func-rename)
```
func Rename(key, newKey K) error
```
Rename изменяет ключ ``key`` на новый ``newKey`` у элемента.

### [func Resize](#func-resize)
```
func Resize(cap int) error
```
Resize изменяет размер кеша.

Если ``cap`` > ёмкости кеша - элементы копируются в новое хранилище.<br>
Если ``cap`` = ёмкости кеша - ничего не меняется.<br>
Елси ``cap`` < текущее количество элементов в кеше - в новое хранилище попадают только наиболее используемые элементы, остальные - урезаются.

### [func SaveJSON](#func-savejson)
```
func SaveJSON(filePath string) error
```
SaveJSON записывает данные из кеша в файл ``filePath`` в формате JSON.

### [func Set](#func-set)
```
func Set(key K, value V) 
```
Set устанавливает новое значение элементу с ключем ``key`` или создает элемент, если его нет.

### [func Values](#func-values)
```
func Values() []V
```
Values возвращает все элементы кеша.