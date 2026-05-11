/*
 * L8Learn Student Player — Configuration
 * Loads login.json and provides API access
 */
(function() {
    'use strict';

    window.PlayerConfig = {
        apiPrefix: '/learn',
        bearerToken: null,
        studentId: null,
        studentName: null,

        async load() {
            const resp = await fetch('/login.json');
            const config = await resp.json();
            this.apiPrefix = config.app.apiPrefix || '/learn';
            this.bearerToken = sessionStorage.getItem('bearerToken');
            this.studentId = sessionStorage.getItem('studentId');
            this.studentName = sessionStorage.getItem('studentName');

            if (!this.bearerToken) {
                window.location.href = '/login.html';
            }
        },

        getHeaders() {
            return {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + this.bearerToken
            };
        },

        resolveEndpoint(path) {
            return this.apiPrefix + path;
        },

        async get(endpoint, query) {
            const body = encodeURIComponent(JSON.stringify({ text: query }));
            const url = this.resolveEndpoint(endpoint) + '?body=' + body;
            const resp = await fetch(url, { method: 'GET', headers: this.getHeaders() });
            return resp.json();
        },

        async post(endpoint, data) {
            const url = this.resolveEndpoint(endpoint);
            const resp = await fetch(url, {
                method: 'POST',
                headers: this.getHeaders(),
                body: JSON.stringify(data)
            });
            return resp.json();
        },

        async put(endpoint, data) {
            const url = this.resolveEndpoint(endpoint);
            const resp = await fetch(url, {
                method: 'PUT',
                headers: this.getHeaders(),
                body: JSON.stringify(data)
            });
            return resp.json();
        }
    };
})();
