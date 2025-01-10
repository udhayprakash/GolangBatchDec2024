# packaging


    calculator
        main.go
        addSub
            addfile.go
            subtractfile.go
        mulDiv
            dividefile.go
            multiplyfile.go
        go.mod


## To create package
	go mod init calculator
	go mod tidy


## For publishing and consuming

	go mod init github.com/username/calculator

	go install github.com/username/calculator

