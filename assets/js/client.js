"use strict";

class Client {
    static async getResponse(url) {
        const response = await fetch(url)

        if (!response.ok) {
            throw new Error('An error has occurred: ${response.status} ')
        }

        return await response.json();
    }
}

function getNamespace() {
    Client.getResponse('http://localhost:8085/v1/k8s/namespaces')
        .then(console.log);
}