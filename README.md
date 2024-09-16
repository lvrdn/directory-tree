# directory-tree
Цель задания - написание программы, выводящую в терминал древовидную структуру директории с указанием файлов и их веса (опция -f). Без указания -f выводятся только папки.

Пример вывода программы с опцией -f для директории testdata:
```
├───project
│       ├───file.txt (19b)
│       └───gopher.png (70372b)
├───static
│       ├───a_lorem
│       │       ├───dolor.txt (empty)
│       │       ├───gopher.png (70372b)
│       │       └───ipsum
│       │               └───gopher.png (70372b)
│       ├───css
│       │       └───body.css (28b)
│       ├───empty.txt (empty)
│       ├───html
│       │       └───index.html (57b)
│       ├───js
│       │       └───site.js (10b)
│       └───z_lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
├───zline
│       ├───empty.txt (empty)
│       └───lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
└───zzfile.txt (empty)
```

Наименование и назначение основных файлов проекта:
1. main.go - разработанный код программы.
2. HW_readme.md - описание условий задания.
3. main_test.go - приложенные тесты к условию задания.
4. папка testdata - приложенные тестовые данные к условию задания.
