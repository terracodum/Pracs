package main

import (
	"testing"
)

// -----------------------------------------------------------------------------
// 1. УНИКАЛЬНЫЕ ЭЛЕМЕНТЫ
// -----------------------------------------------------------------------------
func TestUnique(t *testing.T) {
	input := []string{"apple", "banana", "apple", "cherry", "banana"}
	result := unique(input)

	expected := []string{"apple", "banana", "cherry"}
	if len(result) != len(expected) {
		t.Fatalf("got %v, want %v", result, expected)
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("index %d: got %q, want %q", i, result[i], v)
		}
	}

	// оригинал не должен измениться
	if len(input) != 5 {
		t.Error("original slice was modified")
	}
}

// -----------------------------------------------------------------------------
// 2. ГРУППИРОВКА ПО ПЕРВОЙ БУКВЕ
// -----------------------------------------------------------------------------
func TestGroupByFirstLetter(t *testing.T) {
	input := []string{"apple", "banana", "avocado", "cherry"}
	result := groupByFirstLetter(input)

	check := map[string][]string{
		"a": {"apple", "avocado"},
		"b": {"banana"},
		"c": {"cherry"},
	}

	for key, expected := range check {
		got := result[key]
		if len(got) != len(expected) {
			t.Errorf("key %q: got %v, want %v", key, got, expected)
		}
	}
}

// -----------------------------------------------------------------------------
// 3. ПОДСЧЁТ ВХОЖДЕНИЙ
// -----------------------------------------------------------------------------
func TestWordCount(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b", "a"}
	result := wordCount(input)

	expected := map[string]int{"a": 3, "b": 2, "c": 1}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("key %q: got %d, want %d", k, result[k], v)
		}
	}
}

// -----------------------------------------------------------------------------
// 4. РАЗВЕРНУТЬ СРЕЗ
// -----------------------------------------------------------------------------
func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := reverse(input)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("index %d: got %d, want %d", i, result[i], v)
		}
	}

	// оригинал не должен измениться
	if input[0] != 1 {
		t.Error("original slice was modified")
	}
}

// -----------------------------------------------------------------------------
// 5. ПЕРЕСЕЧЕНИЕ
// -----------------------------------------------------------------------------
func TestIntersect(t *testing.T) {
	result := intersect([]int{1, 2, 3, 4}, []int{2, 4, 6})

	expected := map[int]bool{2: true, 4: true}
	if len(result) != len(expected) {
		t.Errorf("got %v, want [2 4]", result)
	}
	for _, v := range result {
		if !expected[v] {
			t.Errorf("unexpected value %d in result", v)
		}
	}
}

// -----------------------------------------------------------------------------
// 6. ФИЛЬТРАЦИЯ
// -----------------------------------------------------------------------------
func TestFilterActive(t *testing.T) {
	users := []User{
		{Name: "Alice", Active: true},
		{Name: "Bob", Active: false},
		{Name: "Charlie", Active: true},
	}
	result := filterActive(users)

	if len(result) != 2 {
		t.Errorf("got %d users, want 2", len(result))
	}
	for _, u := range result {
		if !u.Active {
			t.Errorf("inactive user %q in result", u.Name)
		}
	}
}

// -----------------------------------------------------------------------------
// 7. СОРТИРОВКА ПО ВОЗРАСТУ
// -----------------------------------------------------------------------------
func TestSortByAge(t *testing.T) {
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 20},
		{Name: "Charlie", Age: 25},
	}
	result := sortByAge(users)

	for i := 1; i < len(result); i++ {
		if result[i].Age < result[i-1].Age {
			t.Errorf("not sorted at index %d: %d > %d", i, result[i-1].Age, result[i].Age)
		}
	}
}

// -----------------------------------------------------------------------------
// 8. МАКСИМУМ ПО ПОЛЮ
// -----------------------------------------------------------------------------
func TestOldest(t *testing.T) {
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 45},
		{Name: "Charlie", Age: 25},
	}
	result, err := oldest(users)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "Bob" {
		t.Errorf("got %q, want Bob", result.Name)
	}

	_, err = oldest([]User{})
	if err == nil {
		t.Error("expected error for empty slice, got nil")
	}
}

// -----------------------------------------------------------------------------
// 9. ПАЛИНДРОМ
// -----------------------------------------------------------------------------
func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"racecar": true,
		"hello":   false,
		"level":   true,
		"world":   false,
		"a":       true,
	}
	for input, expected := range cases {
		if isPalindrome(input) != expected {
			t.Errorf("isPalindrome(%q) = %v, want %v", input, !expected, expected)
		}
	}
}

// -----------------------------------------------------------------------------
// 10. RLE СЖАТИЕ
// -----------------------------------------------------------------------------
func TestCompress(t *testing.T) {
	cases := map[string]string{
		"aaabbc": "a3b2c1",
		"abcd":   "a1b1c1d1",
		"aaa":    "a3",
	}
	for input, expected := range cases {
		if compress(input) != expected {
			t.Errorf("compress(%q) = %q, want %q", input, compress(input), expected)
		}
	}
}

// -----------------------------------------------------------------------------
// 11. СВОЯ ОШИБКА
// -----------------------------------------------------------------------------
func TestValidateAge(t *testing.T) {
	if err := validateAge(25); err != nil {
		t.Errorf("expected nil for valid age, got %v", err)
	}

	err := validateAge(-1)
	if err == nil {
		t.Error("expected error for age -1")
	}

	err = validateAge(200)
	if err == nil {
		t.Error("expected error for age 200")
	}
}

// -----------------------------------------------------------------------------
// 12. ГОРУТИНЫ + КАНАЛ
// -----------------------------------------------------------------------------
func TestRunWorkers(t *testing.T) {
	result := runWorkers()

	if len(result) != 5 {
		t.Errorf("got %d results, want 5", len(result))
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	// 0*2 + 1*2 + 2*2 + 3*2 + 4*2 = 20
	if sum != 20 {
		t.Errorf("got sum %d, want 20", sum)
	}
}

// -----------------------------------------------------------------------------
// 13. ТАЙМАУТ
// -----------------------------------------------------------------------------
func TestWithTimeout(t *testing.T) {
	// канал с данными — должно вернуть значение
	ch := make(chan int, 1)
	ch <- 42
	val, err := withTimeout(ch)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 42 {
		t.Errorf("got %d, want 42", val)
	}

	// пустой канал — должен сработать таймаут
	empty := make(chan int)
	_, err = withTimeout(empty)
	if err == nil {
		t.Error("expected timeout error, got nil")
	}
}

// -----------------------------------------------------------------------------
// 14. WORKER POOL
// -----------------------------------------------------------------------------
func TestWorkerPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	result := workerPool(jobs, 3)

	if len(result) != len(jobs) {
		t.Errorf("got %d results, want %d", len(result), len(jobs))
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	// (1+2+3+4+5) * 2 = 30
	if sum != 30 {
		t.Errorf("got sum %d, want 30", sum)
	}
}
