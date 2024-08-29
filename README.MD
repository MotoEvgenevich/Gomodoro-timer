### English Version

---

# Pomodoro Timer CLI

## Description

The **Pomodoro Timer CLI** is a simple and effective command-line tool designed to help users manage their time using the Pomodoro Technique. This tool allows you to set work intervals (usually 25 minutes) followed by short breaks, helping you stay focused and productive. The timer runs directly in the terminal, providing clear notifications when it's time to take a break or start a new work session. You can customize the length of work periods, breaks, and the number of cycles, making it a flexible tool for various tasks.

## Features

- Set custom work intervals and break durations.
- Run multiple Pomodoro cycles with ease.
- Receive notifications when it's time to work or take a break.
- Simple, easy-to-use interface in the terminal.

## Usage

```bash
go run pomodoro.go --work=25m --break=5m --cycles=4
```

- `--work`: Duration of the work interval (default is 25 minutes).
- `--break`: Duration of the break interval (default is 5 minutes).
- `--cycles`: Number of Pomodoro cycles to run (default is 4).

## Requirements

- Go 1.16 or later

---

### Russian Version

---

# CLI-таймер для Помидоро

## Описание

**CLI-таймер для Помидоро** — это простой и эффективный инструмент командной строки, разработанный для управления временем с использованием техники Помидоро. Этот таймер позволяет установить интервалы работы (обычно 25 минут), за которыми следуют короткие перерывы, помогая сохранять концентрацию и продуктивность. Таймер работает прямо в терминале, предоставляя чёткие уведомления о начале нового рабочего периода или перерыва. Вы можете настроить длительность рабочих периодов, перерывов и количество циклов, что делает его гибким инструментом для выполнения различных задач.

## Функции

- Установка пользовательских рабочих интервалов и длительности перерывов.
- Запуск нескольких циклов Помидоро.
- Уведомления о начале работы или перерыва.
- Простой и удобный интерфейс в терминале.

## Использование

```bash
go run pomodoro.go --work=25m --break=5m --cycles=4
```

- `--work`: Длительность рабочего интервала (по умолчанию 25 минут).
- `--break`: Длительность перерыва (по умолчанию 5 минут).
- `--cycles`: Количество циклов Помидоро (по умолчанию 4).

## Требования

- Go 1.16 или новее
