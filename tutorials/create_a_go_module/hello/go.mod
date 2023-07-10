module github.com/Yashikab/go_practice/tutorials/create_a_go_module/hello

go 1.20

replace github.com/Yashikab/go_practice/tutorials/create_a_go_module/greetings => ../greetings

require github.com/Yashikab/go_practice/tutorials/create_a_go_module/greetings v0.0.0-00010101000000-000000000000
