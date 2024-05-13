#!/bin/bash

# Установка PostgreSQL
sudo apt-get update
sudo apt-get install postgresql

# Запуск сервера PostgreSQL
sudo service postgresql start

# Создание нового пользователя базы данных
sudo -u postgres createuser --superuser $USER

# Вход в интерактивный режим psql
psql -U postgres