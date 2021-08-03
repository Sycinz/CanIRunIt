package main

import (
	"bufio"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
	"net"
	"os"
	"strings"
)

type games struct {
	ID int
	name string;
	system string;
	cpu string;
	gpu string;
	ram int;
	graphics string;
	fps int;
}

func main() {
	listener, err := net.Listen("tcp", ":50153");

	if err != nil {
		os.Exit(1);
	}
	defer listener.Close();

	for ;; {
			connect, err := listener.Accept();
			if err != nil {
				os.Exit(1);
			}
			go handler(connect);
	}
}

func handler(connection net.Conn) {
	buff, err := bufio.NewReader(connection).ReadString('\n');
	if err != nil {
		fmt.Println("Wystąpił błąd przy zczytywaniu zawartości");
	}
	splitted := strings.Split(string(buff), ",");
	fmt.Printf("%v\n", buff);
	fmt.Printf("%v\n", splitted[0]);

	switch splitted[0] {
		case "add":
			db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
			//defer db.Close();

			if err != nil {
				fmt.Println("Błąd otwierania bazy danych");
			}
			//var tag Tag;

			var nazwa games;

			db.Last(&nazwa); //SELECT id FROM games ORDER BY id DESC LIMIT 1
			fmt.Printf("ID: %v", nazwa.ID);
			//db.Query("INSERT INTO games VALUES (?,?,?,?,?,?,?,?)", splitted[1], splitted[2], splitted[3], splitted[4], splitted[5], splitted[6], splitted[7]);
			// &tag.ID, &tag.name, &tag.system, &tag.cpu, &tag.gpu, &tag.ram, &tag.graphics, &tag.fps

			if err != nil {
				fmt.Println("Błąd implementacji do bazy danych"); // implementacji do bazy danych
			}
			fmt.Println(&nazwa);

	}
	connection.Write([]byte("Test"));
	connection.Close();
}
