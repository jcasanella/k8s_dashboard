"use strict";

class Client {
    static async getResponse(endpoint, apiVersion = 'v1', apiType = 'k8s') {
        let serverAddress = 'http://localhost:8085';
        const response = await fetch(serverAddress + '/' + apiVersion + '/' + apiType + '/' + endpoint)

        if (!response.ok) {
            throw new Error('An error has occurred: ${response.status} ')
        }

        return await response.json();
    }
}

function getNamespace() {
    Client.getResponse('namespaces')
        .then(processJson);
}

function getPods() {
    Client.getResponse('pods')
        .then(processJson);
}

function getConfigMaps() {
    Client.getResponse('configmaps')
        .then(processJson);
}

function countPods() {
    Client.getResponse('pods/count')
        .then(console.log);
}

function countConfigMaps() {
    Client.getResponse('configmaps/count')
        .then(console.log);
}

function processJson(value) {
    const jsonString = JSON.stringify(value);
    const jsonObj = JSON.parse(jsonString);
    Array.from(jsonObj).forEach(element => console.log(element));
}