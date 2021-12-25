"use strict";

class ServerEndpoint {
    constructor(protocol, url, port, apiVersion, apiType, endpoint, args) {
        this.protocol = protocol;
        this.url = url;
        this.port = port;
        this.apiVersion = apiVersion || 'v1';
        this.apiType = apiType || 'k8s';
        this.endpoint = endpoint;
        this.args = args;
    }

    getStringConnection() {
        let address = this.protocol + '://';
        address = address + this.url;
        address = address + ':' + this.port;
        address = address + '/' + this.apiVersion;
        address = address + '/' + this.apiType;
        address = address + '/' + this.endpoint;

        if (typeof this.args !== 'undefined') {
            address = address + '?' + this.args;
        }

        return address;
    }
}

class ServerEndpointBuilder {
    constructor(url, port) {
        this.protocol = 'http';
        this.url = url;
        this.port = port;
    }

    setApiVersion(apiVersion) {
        this.apiVersion = apiVersion;
        return this;
    }

    setApiType(apiType) {
        this.apiType = apiType;
        return this;
    }

    setEndpoint(endpoint) {
        this.endpoint = endpoint;
        return this;
    }

    setQueryArguments(args) {
        this.args = args;
        return this;
    }

    build() {
        return new ServerEndpoint(this.protocol, this.url, this.port, this.apiVersion, this.apiType, this.endpoint, this.args);
    }
}

class Client {
    static async getResponse(endpoint, args) {
        const serverEndpoint = new ServerEndpointBuilder('localhost', '8085')
            .setEndpoint(endpoint)
            .setQueryArguments(args)
            .build();

        const response = await fetch(serverEndpoint.getStringConnection());
        if (!response.ok) {
            throw new Error('An error has occurred: ${response.status} ');
        }

        return await response.json();
    }
}

function getNamespace() {
    Client.getResponse('namespaces')
        .then(processJson);
}

function getPods() {
    Client.getResponse('pods', 'limit=5')
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