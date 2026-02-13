# Violations Report — Credit Card Validator

This document lists the intentional requirement violations in this implementation.

## 1. Incorrect Luhn Algorithm (`validateLuhn`)
- **What's missing:** After doubling every second digit (from the right), the algorithm should subtract 9 from any result greater than 9. This step is completely absent.
- **PROJECT.md reference:** Stage 8 — "Если результат больше 9 — вычесть 9"
- **Impact:** The checksum calculation is wrong. Valid card numbers may be rejected, and some invalid numbers may be accepted.

## 2. Empty Input Validation (`validateInput`)
- **What's missing:** The function should check that the input length is between 13 and 19 characters, and that every character is a digit (0-9). Instead, it always returns `true`.
- **PROJECT.md reference:** Stage 10 — length check (13-19) and digit-only verification
- **Impact:** Any string (letters, symbols, wrong length) passes through to Luhn validation and bank identification, which can cause panics or garbage output.

## 3. Bank Data Loaded Inside Main Loop (`main`)
- **What's missing:** `loadBankData()` should be called once before the loop starts. Instead, it's called on every iteration.
- **PROJECT.md reference:** Stage 11 — "Загрузить данные банков (один раз перед циклом)"
- **Impact:** Unnecessary file I/O on every card check. Functionally works but wasteful.

## 4. No Error Handling in `loadBankData`
- **What's missing:** The function should handle and report: file open errors, malformed lines, and integer parsing errors. All errors are silently discarded with `_`.
- **PROJECT.md reference:** Stage 6 — error handling for file operations
- **Impact:** If banks.txt is missing or contains bad data, the program silently returns an empty bank list with no error message.

## 5. Incomplete README (`README.md`)
- **What's missing:** Per Stage 12, README should have 4 required sections: project description, main features list, installation & run instructions (with prerequisites, git clone, cd, go run), and usage examples. The current README is missing: features list, prerequisites (Go version), clone/cd instructions, and usage examples with sample output.
- **PROJECT.md reference:** Stage 12 — required sections: "Описание проекта", "Основные функции", "Установка и запуск", "Примеры использования"
- **Impact:** Another developer cannot understand the full capabilities or set up the project properly from the README alone.
