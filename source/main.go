package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type identity struct {
	typeID string
	noID   string
}

type loker struct {
	noLoker  int
	identity identity
}

var listLoker []loker

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------------------")
	fmt.Println("|---- Loker Go -----|")
	fmt.Println("---------------------")
	fmt.Println("Silahkan tentukan jumlah loker...")
	fmt.Println("Contoh: init 5")
	for {
		if len(listLoker) < 1 {
			fmt.Print("0 > ")
		} else {
			fmt.Print("1 > ")
		}
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		command := strings.Split(text, " ")
		isInit := checkInitLoker(command)
		if !isInit {
			break
		}
		switch strings.ToLower(command[0]) {
		case "init":
			if len(command) != 2 {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			if command[1] == "" {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			initLoker(command[0], command[1])
		case "status":
			statusLoker()
		case "input":
			if len(command) != 3 {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			if command[1] == "" {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			inputLoker(command[1], command[2])
		case "leave":
			if len(command) != 2 {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			if command[1] == "" {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			leaveLoker(command[1])
		case "find":
			if len(command) != 2 {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			if command[1] == "" {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			findLoker(command[1])
		case "search":
			if len(command) != 2 {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			if command[1] == "" {
				fmt.Println("Peringatan: Perintah salah!")
				break
			}
			searchLoker(command[1])
		case "exit":
			exitLoker()
		case "clear":
			clear()
		case "":
			fmt.Println("Peringatan: Perintah tidak boleh kosong!")
		default:
			fmt.Println("Peringatan: Perintah tidak dikenali!")
		}
	}
}

func checkInitLoker(command []string) bool {
	if len(listLoker) < 1 && strings.ToLower(command[0]) != "init" && strings.ToLower(command[0]) != "exit" && strings.ToLower(command[0]) != "clear" {
		println("Peringatan: Tentukan jumlah loker terlebih dahulu!")
		return false
	}

	return true
}

func getNoLoker(noLokerStr string) (int, bool) {
	noLoker, err := strconv.Atoi(noLokerStr)
	if err != nil {
		fmt.Println("Peringatan: Jumlah loker harus bernilai integer!")
		return 0, false
	}
	if noLoker < 1 {
		fmt.Println("Peringatan: Jumlah loker harus lebih dari 0!")
		return 0, false
	}
	return noLoker, true
}

func initLoker(command string, totalLoker string) bool {
	if len(listLoker) > 1 && command == "init" {
		println("Peringatan: Jumlah loker sudah di tentukan!")
		return false
	}
	firstInt := string([]rune(totalLoker)[0])
	if firstInt == "0" {
		fmt.Println("Peringatan: Nilai loker harus di awali lebih dari 0!")
		return false
	}
	noLoker, err := getNoLoker(totalLoker)
	if !err {
		return false
	}
	for i := 0; i < noLoker; i++ {
		objLoker := loker{
			noLoker: i + 1,
		}
		listLoker = append(listLoker, objLoker)
	}
	fmt.Println("Berhasil membuat loker dengan jumlah", totalLoker)
	return true
}

func statusLoker() {
	fmt.Println("| No Loker | Tipe Identitas | No Identitas |")
	for i := 0; i < len(listLoker); i++ {
		ti := ""
		ni := ""
		if (identity{}) != listLoker[i].identity {
			ti = listLoker[i].identity.typeID
			ni = listLoker[i].identity.noID
		}
		fmt.Println("| ", listLoker[i].noLoker, "      | ", ti, "             | ", ni, "           |")
	}
}

func inputLoker(typeID string, noID string) bool {
	isFull := false
	for i := 0; i < len(listLoker); i++ {
		if (identity{}) == listLoker[i].identity {
			listLoker[i].identity.typeID = strings.ToUpper(typeID)
			listLoker[i].identity.noID = noID
			fmt.Println("Kartu Identitas tersimpan di loker nomor ", i+1)
			break
		} else if i+1 == len(listLoker) {
			isFull = true
		}
	}
	if isFull {
		fmt.Println("Peringatan: Maaf loker sudah penuh!")
		return false
	}
	return true
}

func leaveLoker(noLokerStr string) bool {
	noLoker, err := getNoLoker(noLokerStr)
	if !err {
		return false
	}
	if listLoker[noLoker-1].identity == (identity{}) {
		fmt.Println("Loker nomor", noLoker, "kosong")
		return false
	}
	listLoker[noLoker-1].identity = identity{}
	fmt.Println("Loker nomor ", noLoker, "berhasil di kosongkan")
	return true
}

func findLoker(noID string) {
	isExists := false
	index := 0
	for i := 0; i < len(listLoker); i++ {
		if (identity{}) != listLoker[i].identity {
			if listLoker[i].identity.noID == noID {
				isExists = true
				index = i
				break
			}
		}
	}
	if isExists {
		fmt.Println("Kartu identitas tersebut berada di loker nomor", index+1)
	} else {
		fmt.Println("Nomor identitas tidak ditemukan")
	}
}

func searchLoker(typeID string) {
	listResult := ""
	for i := 0; i < len(listLoker); i++ {
		if (identity{}) != listLoker[i].identity {
			if listLoker[i].identity.typeID == strings.ToUpper(typeID) {
				if listResult == "" {
					listResult = listLoker[i].identity.noID
				} else {
					listResult = listResult + "," + listLoker[i].identity.noID
				}
			}
		}
	}
	fmt.Println(listResult)
}

func exitLoker() {
	fmt.Println("Sampai jumpa!")
	os.Exit(0)
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
