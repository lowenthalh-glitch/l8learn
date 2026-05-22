(function() {
    'use strict';

    var USERS_ENDPOINT = '/73/users';
    var DEFAULT_PASSWORD = '12345678';

    function getHeaders() {
        return Object.assign({ 'Content-Type': 'application/json' },
            typeof getAuthHeaders === 'function' ? getAuthHeaders() : {});
    }

    async function postUser(user) {
        try {
            var resp = await fetch(Layer8DConfig.resolveEndpoint(USERS_ENDPOINT), {
                method: 'POST', headers: getHeaders(), body: JSON.stringify(user)
            });
            if (!resp.ok) {
                console.warn('[UserProvisioning] user creation returned', resp.status, await resp.text());
                return false;
            }
            return true;
        } catch (e) {
            console.warn('[UserProvisioning] user creation failed', e);
            return false;
        }
    }

    async function createStudentUser(student) {
        var studentId = student.studentId;
        if (!studentId) return;
        var email = (student.firstName + '.' + student.lastName + '@student.l8learn.local').toLowerCase();
        var roles = {}; roles['student'] = true;
        var ok = await postUser({
            userId: studentId, fullName: (student.firstName + ' ' + student.lastName).trim(),
            email: email, accountStatus: 'ACCOUNT_STATUS_ACTIVE',
            portal: 'app.html', password: { hash: DEFAULT_PASSWORD }, roles: roles
        });
        if (ok) Layer8DNotification.success('Student user account created');
        else    Layer8DNotification.warning('Student saved but user account creation failed');
    }

    async function createGuardianUser(guardian) {
        var email = guardian.email;
        if (!email) return;
        var roles = {}; roles['guardian'] = true;
        var ok = await postUser({
            userId: guardian.guardianId, fullName: (guardian.firstName + ' ' + guardian.lastName).trim(),
            email: email, accountStatus: 'ACCOUNT_STATUS_ACTIVE',
            portal: 'app.html', password: { hash: DEFAULT_PASSWORD }, roles: roles
        });
        if (ok) Layer8DNotification.success('Guardian user account "' + email + '" created');
        else    Layer8DNotification.warning('Guardian saved but user account creation failed');
    }

    async function createTeacherUser(teacher) {
        var email = teacher.email;
        if (!email) return;
        var roles = {}; roles['teacher'] = true;
        var ok = await postUser({
            userId: teacher.teacherId, fullName: (teacher.firstName + ' ' + teacher.lastName).trim(),
            email: email, accountStatus: 'ACCOUNT_STATUS_ACTIVE',
            portal: 'app.html', password: { hash: DEFAULT_PASSWORD }, roles: roles
        });
        if (ok) Layer8DNotification.success('Teacher user account "' + email + '" created');
        else    Layer8DNotification.warning('Teacher saved but user account creation failed');
    }

    window.LearnUserProvisioning = {
        createStudentUser: createStudentUser,
        createGuardianUser: createGuardianUser,
        createTeacherUser: createTeacherUser
    };

    // File viewer — fetches encrypted file from FileStore and shows text in a popup
    window.LearnFileViewer = {
        view: function(storagePath, title) {
            if (!storagePath) return;
            var endpoint = Layer8DConfig.resolveEndpoint('/0/FileStore');
            fetch(endpoint, {
                method: 'PUT',
                headers: getHeaders(),
                body: JSON.stringify({ storagePath: storagePath })
            }).then(function(resp) {
                if (!resp.ok) throw new Error('Failed to fetch file');
                return resp.json();
            }).then(function(data) {
                if (!data.fileData) throw new Error('No file data');
                var text = atob(data.fileData);
                var escaped = Layer8DUtils.escapeHtml(text);
                Layer8DPopup.show({
                    title: title || 'Evaluation Document',
                    content: '<pre style="white-space:pre-wrap;word-wrap:break-word;font-size:13px;line-height:1.6;padding:12px;max-height:70vh;overflow-y:auto;">' + escaped + '</pre>',
                    size: 'xlarge'
                });
            }).catch(function(err) {
                Layer8DNotification.error('Failed to view file: ' + err.message);
            });
        }
    };

    // Override the default file download button to open in browser instead
    if (window.Layer8DFormsFields && Layer8DFormsFields.onFileDownload) {
        Layer8DFormsFields.onFileDownload = function(btn) {
            var path = btn.getAttribute('data-storage-path');
            var name = btn.getAttribute('data-file-name') || 'Evaluation Document';
            LearnFileViewer.view(path, name);
        };
    }
})();
