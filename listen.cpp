#include <cstring>
#include <iostream>
#include <netinet/in.h>
#include <sys/socket.h>
#include <unistd.h>

int listen(std::string IPv4port)
{
    // creating socket
    int serverSocket = socket(AF_INET, SOCK_STREAM, 0);

    // specifying the address
    sockaddr_in serverAddress;
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_port = htons(std::stoi(IPv4port));
    serverAddress.sin_addr.s_addr = INADDR_ANY;

    // binding socket.
    int bindedSocketStatus = bind(serverSocket, (struct sockaddr *)&serverAddress,
                                  sizeof(serverAddress));

    if (bindedSocketStatus == 0)
    {
        std::cout << "Binded Successfully\n";
    }
    else
    {
        throw std::invalid_argument("Port " + IPv4port + " sucks at being a binded to\n");
    }

    // listening to the assigned socket
    listen(serverSocket, 5);

    while (true)
    {

        // accepting connection request
        int clientSocket;

        // recieving data
        char buffer[4096];

        clientSocket = accept(serverSocket, nullptr, nullptr);
        buffer[4096] = {0};
        recv(clientSocket, buffer, sizeof(buffer), 0);
        std::cout << "Message from client: " << buffer
                  << std::endl;
        const char *message = "HTTP/1.1 200 OK\nContent-Type: text/html; charset=UTF-8\n\nHello bean-brain\n";
        send(clientSocket, message, strlen(message), 0);
        close(clientSocket);
    }

    close(serverSocket);

    return 0;
}