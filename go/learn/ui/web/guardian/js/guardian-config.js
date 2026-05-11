/*
 * L8Learn Guardian Portal — Configuration & Auth
 */
(function() {
    'use strict';

    window.GuardianConfig = {
        apiPrefix: '/learn',
        bearerToken: null,
        guardianId: null,
        familyId: null,
        studentIds: [],

        async load() {
            var resp = await fetch('/login.json');
            var config = await resp.json();
            this.apiPrefix = config.app.apiPrefix || '/learn';
            this.bearerToken = sessionStorage.getItem('bearerToken');
            this.guardianId = sessionStorage.getItem('guardianId');
            this.familyId = sessionStorage.getItem('familyId');
            var ids = sessionStorage.getItem('studentIds');
            this.studentIds = ids ? JSON.parse(ids) : [];

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

        async get(endpoint, query) {
            var body = encodeURIComponent(JSON.stringify({ text: query }));
            var url = this.apiPrefix + endpoint + '?body=' + body;
            var resp = await fetch(url, { method: 'GET', headers: this.getHeaders() });
            return resp.json();
        },

        async post(endpoint, data) {
            var url = this.apiPrefix + endpoint;
            var resp = await fetch(url, { method: 'POST', headers: this.getHeaders(), body: JSON.stringify(data) });
            return resp.json();
        }
    };
})();
