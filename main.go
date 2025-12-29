package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// repoPath: hozirgi ishchi papka
	repoPath, _ := os.Getwd()

	for {
		filename := fmt.Sprintf("%s/hello_%d.json", repoPath, time.Now().Unix())
		content := `{"message": "Hello World"}`

		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Fayl yaratishda xato:", err)
			continue
		}

		cmdAdd := exec.Command("git", "-C", repoPath, "add", ".")
		cmdAdd.Run()

		cmdCommit := exec.Command("git", "-C", repoPath, "commit", "-m", fmt.Sprintf("Auto update %s", time.Now().Format(time.RFC3339)))
		cmdCommit.Run()

		cmdPush := exec.Command("git", "-C", repoPath, "push", "origin", "main")
		cmdPush.Run()

		fmt.Println("Fayl push qilindi:", filename)

		time.Sleep(12 * time.Hour) // 12 soat kutadi
	}
}
