 Web Services



 D1         WEb Services     D2
    ----(http/tcp/udp)--------

D1      (web application)    Human



Device 
    ip address (recognizes device in the internet)
        Decided by Network (NOT FIXED)
    MAC address ( recognizses the physical address of server)
        Decided by manufacturer (Mostly Fixed )
    port address (recognizes the application process, within the server)
        Ephimerical ports (common for all -- fixed)
        Non-Ephimerical - specific to applictaions ( decided by OS/appliction)


ip address 
    1) ip v4 -- 32 bit 
                (8bits . 8 bits . 8 bits . 8 bits)/ 
                (1 byte. 1byte. 1byte .     1byte)
                ex:  192.168.1.1
    2) ip v6  - 128 bit
             - Ex: 2001:0db8:85a3:0000:0000:8a2e:0370:7334

ip v4 
    Public addresses
    Private addresses
    Experimental/ fix local
        127.0.0.1  -- LoopBack address

    classes 
        class A 
            Range       1.0.0.0 to 126.255.255.255
            FirstOctet  1 to 126
            subnetMask  255.0.0.0 or /8
            size        Large networks (16,777,214 hosts per network)
            ex:         10.0.0.1
        class B
            Range       128.0.0.0 to 191.255.255.255
            FirstOctet  128 to 191
            subnetMask  255.255.0.0 or /16
            size         Medium-sized networks (65,534 hosts per network)
            ex:         172.16.0.1
        class C
            Range       192.0.0.0 to 223.255.255.255
            FirstOctet  192 to 223
            subnetMask  255.255.255.0 or /24
            size         Small-sized networks (254 hosts per network)
            ex:         192.168.1.1
        class D
            Range       224.0.0.0 to 239.255.255.255
            FirstOctet  224 to 239
            size        Used for multicast groups, not for individual devices.
            Example:    224.0.0.1 (used for all hosts on a network).
        class E
            Range       224.0.0.0 to 239.255.255.255
            FirstOctet  240 to 255
            size        Reserved for experimental use.
                        Not used in pubic or private networks
            Ex:         240.0.0.1

Private IP Address Ranges
========

-------------------------------------------------------------
Class	Private IP Range	                Subnet Mask
-------------------------------------------------------------
A	    10.0.0.0 to 10.255.255.255	        255.0.0.0 or /8
B	    172.16.0.0 to 172.31.255.255	    255.240.0.0 or /12
C	    192.168.0.0 to 192.168.255.255	    255.255.0.0 or /16
-------------------------------------------------------------


Special IP Addresses
===
    
    Loopback Address: 127.0.0.1 (used for testing on the local machine).
    APIPA (Automatic Private IP Addressing): 169.254.0.0 to 169.254.255.255 (used when a device cannot obtain an IP address from a DHCP server).

    Broadcast Address: 255.255.255.255 (used to send data to all devices on a network).


Subnetting and CIDR
====

Subnetting:
    
    Dividing a network into smaller sub-networks to improve efficiency and security
    
        application  (public subnet)
        database      (private subnet)

CIDR (Classless Inter-Domain Routing)

    A method for allocating IP addresses and routing IP packets more efficiently.
    Represented as IP_address/Prefix_length (e.g., 192.168.1.0/24).  (192.168.1.0 to 192.168.1.255)


