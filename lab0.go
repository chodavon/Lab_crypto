package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Compute all hashes for a string
func hash(s string) map[string]string {
	return map[string]string{
		"MD5":      fmt.Sprintf("%x", md5.Sum([]byte(s))),
		"SHA1":     fmt.Sprintf("%x", sha1.Sum([]byte(s))),
		"SHA256":   fmt.Sprintf("%x", sha256.Sum256([]byte(s))),
		"SHA512":   fmt.Sprintf("%x", sha512.Sum512([]byte(s))),
		"SHA3-256": fmt.Sprintf("%x", sha3.Sum256([]byte(s))),
	}
}

// local helper to compute SHA-512 hex (used instead of external crack.Sha512Hex)
func Sha512Hex(pass string) string {
	sum := sha512.Sum512([]byte(pass))
	return fmt.Sprintf("%x", sum)
}

// Simple interactive hash comparison task
func task() {
	var input1, input2 string

	fmt.Println("========== Task 0 ==========")
	fmt.Print("Enter value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Enter value 2: ")
	fmt.Scanln(&input2)

	hash1 := hash(input1)
	hash2 := hash(input2)

	fmt.Println("\n======= Hash Comparison =======")
	for _, algo := range []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3-256"} {
		result := "No Match!"
		if hash1[algo] == hash2[algo] {
			result = "Match!"
		}
		fmt.Printf("%s => %s\n", algo, result)
	}
}

// CrackMD5 tries each word in the wordlist and prints each attempt (verbose).
func CrackMD5(target, wordlistPath string) (string, error) {
	target = strings.ToLower(strings.TrimSpace(target))

	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", fmt.Errorf("open wordlist: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	attempt := 0
	for sc.Scan() {
		attempt++
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// Verbose: print each attempt
		fmt.Printf("Try #%d: %s\n", attempt, line)

		sum := md5.Sum([]byte(line))
		if hex.EncodeToString(sum[:]) == target {
			fmt.Printf("FOUND -> #%d : %s  (md5=%s)\n", attempt, line, hex.EncodeToString(sum[:]))
			return line, nil
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("scan error: %w", err)
	}
	return "", nil
}

// Task4: decode or verify the cat-CTF flag
func task4() {
	pattern := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`
	expected := "cryptoCTF{meowmeow}"

	fmt.Println("========== Task 4 ==========")
	fmt.Print("Paste regex or flag: ")
	in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	in = strings.TrimSpace(in)

	// If user pasted the regex literally, print decoded flag
	if in == pattern {
		fmt.Println("Decoded flag:", expected)
		return
	}

	// If user pasted the flag, check it
	if in == expected {
		fmt.Println("Valid flag ")
		return
	}

	// Otherwise, try matching the regex against the input
	re := regexp.MustCompile(pattern)
	if re.MatchString(in) {
		fmt.Println("Valid flag ")
	} else {
		fmt.Println("Invalid flag  (expected: cryptoCTF{meowmeow})")
	}
}

func task2() {
	fmt.Println("========== Task 2 ==========")
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9" // hash to crack
	wordlist := "nord_vpn.txt"                           // put wordlist here

	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalf("open wordlist: %v", err)
	}
	defer f.Close()

	out, err := os.Create("verbose.txt") // verbose log file
	if err != nil {
		log.Fatalf("create verbose log: %v", err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		count++
		pass := strings.TrimSpace(scanner.Text())

		// compute SHA1 hex
		sum := sha1.Sum([]byte(pass))
		h := fmt.Sprintf("%x", sum)

		// save attempt
		if _, err := fmt.Fprintf(out, "#%d %s -> %s\n", count, pass, h); err != nil {
			log.Printf("warning: could not write verbose log: %v", err)
		}

		if h == target {
			fmt.Println("FOUND:", pass)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}
	fmt.Println("Password not found")
}

func taskSHA512() {
	fmt.Println("========== Task SHA512 =========")
	rawTarget := `485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f0 26ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38`
	target := strings.ReplaceAll(strings.ReplaceAll(rawTarget, " ", ""), "\n", "")
	target = strings.TrimSpace(target)

	wordlist := "nord_vpn.txt"

	// open wordlist
	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalf("open wordlist: %v", err)
	}
	defer f.Close()

	// create verbose file
	out, err := os.Create("verbose.txt")
	if err != nil {
		log.Fatalf("create verbose log: %v", err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)

	// increase buffer in case of very long lines
	const maxCapacity = 10 * 1024 * 1024 // 10MB
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	count := 0
	for scanner.Scan() {
		count++
		pass := strings.TrimSpace(scanner.Text())
		if pass == "" {
			continue
		}

		// use local helper
		h := Sha512Hex(pass)

		if _, err := fmt.Fprintf(out, "#%d %s -> %s\n", count, pass, h); err != nil {
			log.Printf("warning: couldn't write to verbose.txt: %v", err)
		}

		if h == target {
			fmt.Println("FOUND:", pass)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}

	fmt.Println("Password not found")
}

func main() {
	// Task 0: hash comparison
	task()

	// Task 1: MD5 cracking
	// Replace "nord_vpn.txt" with the correct path to your downloaded wordlist file.
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlist := "nord_vpn.txt"

	fmt.Println("\n========== Task 1: MD5 Crack ==========")
	found, err := CrackMD5(targetHash, wordlist)
	if err != nil {
		fmt.Println("Error:", err)
	} else if found != "" {
		fmt.Println("\nResult: Password found ->", found)
	} else {
		fmt.Println("\nResult: Password NOT found in the wordlist.")
	}

	task2()
	taskSHA512()
	// Task 4: CTF flag verify/decode
	fmt.Println()
	task4()
}
