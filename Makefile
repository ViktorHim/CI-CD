EXEC=usr/bin/latin_checker
SRC=src/main.go
GO := $(shell which go)

ifeq ($(GO),)
$(error "Go не установлен! Установите golang для сборки.")
endif

DEB_NAME=latin-checker
DEB_VERSION=1.0
DEB_DIR=deb_package

all: $(EXEC)

$(EXEC): $(SRC) | usr/bin
	@echo "Сборка программы..."
	go build -o $(EXEC) $(SRC)
	@echo "Сборка завершена."

usr/bin:
	@mkdir -p usr/bin

deb: deb_structure
	@echo "Создание deb-пакета..."
	mkdir -p $(DEB_DIR)/usr/local/bin
	cp $(EXEC) $(DEB_DIR)/usr/local/bin/
	@printf "Package: $(DEB_NAME)\nVersion: $(DEB_VERSION)\nSection: utils\nPriority: optional\nArchitecture: amd64\nDepends: golang-go (>=1.20)\nMaintainer: Viktor <vhimlihmail@gmail.com>\nDescription: Проверка латинских квадратов\n Простая программа на Go.\n" > $(DEB_DIR)/DEBIAN/control
	dpkg-deb --build $(DEB_DIR) $(DEB_NAME)-$(DEB_VERSION).deb
	@echo "deb-пакет создан: $(DEB_NAME)-$(DEB_VERSION).deb"

deb_structure:
	@mkdir -p $(DEB_DIR)/DEBIAN
	@mkdir -p $(DEB_DIR)/usr/local/bin

clean:
	@echo "Очистка..."
	rm -rf usr/
	rm -rf $(DEB_DIR)
	rm -f *.deb
	@echo "Очистка завершена."

.PHONY: all clean deb deb_structure