package main

func main() {
	// lis, err := net.Listen("tcp", ":9000")
	// println(lis)
	// if err != nil {
	// 	log.Printf("Error listening: %s", err.Error())
	// 	os.Exit(1)
	// }

	c := &b.generatePasswords.passwords.generatepasswords.GeneratePasswords{
		Length: 8,
	}

	println(c.CreatePassword())
}
