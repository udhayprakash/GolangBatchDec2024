# BuildTags


    //go:build   (from Go 1.17)
    // +build

NOTE: Ensure there is one whitespace after the build tag declaration


1) Single Tag


    //go:build tag_name
    // +build tag_name

2) Multiple Tags 

    AND Condition

            //go:build tag1 && tag2
            // +build tag1,tag2

    OR Condition

            //go:build tag1 || tag2
            // +build tag1 tag2

3) Negation


            //go:build !tag_name
            // +build !tag_name
            

### Platform Specific Build 

    To run only on linux
        //go:build linux
        // +build linux

    To exclude specific platform

        //go:build !windows
        // +build !windows


Environment Specific Builds


            //go:build test
            // +build test


Hybrid 


        //go:build (linux || darwin) && amd64
        // +build linux darwin,amd64


Architecture-Specific Builds


            //go:build amd64
            // +build amd64


Debug tags 


            //go:build debug
            // +build debug


All tags 
========
OS Tags         : linux, windows, darwin, freebsd, etc.
Architecture Tags: amd64, arm, arm64, 386, etc.
Build Modes     : cgo, gc, purego, etc.



Set the Build Tags During Compilation
=======================================

// Build with a specific tag
go build -tags debug

// Run tests with a specific tag
go test -tags test

//  Combine multiple tags
go build -tags "debug linux"
