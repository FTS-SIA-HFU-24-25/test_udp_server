import socket

HOST = "192.168.1.50"
PORT = 3000

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind((HOST, PORT))

print("Server is running!")

while True:
    buffer, addr = sock.recvfrom(1024)
    print(len(buffer), addr)
    print(buffer)
