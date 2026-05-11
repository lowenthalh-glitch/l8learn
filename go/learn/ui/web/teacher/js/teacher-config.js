/*
 * L8Learn Teacher Portal — Configuration & Auth
 */
(function() {
    'use strict';

    window.TeacherConfig = {
        apiPrefix: '/learn',
        bearerToken: null,
        teacherId: null,
        classroomIds: [],

        async load() {
            var resp = await fetch('/login.json');
            var config = await resp.json();
            this.apiPrefix = config.app.apiPrefix || '/learn';
            this.bearerToken = sessionStorage.getItem('bearerToken');
            this.teacherId = sessionStorage.getItem('teacherId');
            var ids = sessionStorage.getItem('classroomIds');
            this.classroomIds = ids ? JSON.parse(ids) : [];

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
