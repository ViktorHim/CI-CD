EXEC=build/latin_checker
SRC=src/main.go
GO := $(shell which go)

ifeq ($(GO),)
$(error "Go не установлен! Установите golang для сборки.")
endif

DEB_NAME=latin-checker
DEB_VERSION=1.0
DEB_DIR=deb_package

all: $(EXEC)

$(EXEC): $(SRC) | build
	@echo "Сборка программы..."
	go build -o $(EXEC) $(SRC)
	@echo "Сборка завершена."

build:
	@mkdir -p build
	@echo "Папка build создана."

deb: $(EXEC) | deb_structure
	@echo "Создание deb-пакета..."
	mkdir -p $(DEB_DIR)/usr/local/bin
	cp $(EXEC) $(DEB_DIR)/usr/local/bin/
	@printf "Package: $(DEB_NAME)\nVersion: $(DEB_VERSION)\nSection: utils\nPriority: optional\nArchitecture: amd64\nDepends: golang-go (>=1.20)\nMaintainer: Viktor <vhimlihmail@gmail.com>\nDescription: Проверка латинских квадратов\n Простая программа на Go, проверяющая, является ли заданная матрица латинским квадратом.\n" > $(DEB_DIR)/DEBIAN/control
	dpkg-deb --build $(DEB_DIR)
	@echo "deb-пакет создан: $(DEB_DIR).deb"

deb_structure:
	@mkdir -p $(DEB_DIR)/DEBIAN
	@echo "Структура deb-пакета создана."

clean:
	@echo "Очистка..."
	rm -rf build
	rm -rf $(DEB_DIR)
	@echo "Папки build и deb_package удалены."

.PHONY: all clean build deb deb_structure
