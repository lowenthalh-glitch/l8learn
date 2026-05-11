/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 * Licensed under the Apache License, Version 2.0
 */
function getAuthHeaders() {
    var bearerToken = sessionStorage.getItem('bearerToken');
    return {
        'Authorization': bearerToken ? 'Bearer ' + bearerToken : '',
        'Content-Type': 'application/json'
    };
}

function logout() {
    sessionStorage.removeItem('bearerToken');
    localStorage.removeItem('bearerToken');
    window.location.href = 'l8ui/login/index.html';
}

(function() {
    'use strict';
    document.addEventListener('DOMContentLoaded', async function() {
        var bearerToken = sessionStorage.getItem('bearerToken');
        if (bearerToken) {
            localStorage.setItem('bearerToken', bearerToken);
            window.bearerToken = bearerToken;
        }

        if (window.Layer8DConfig && Layer8DConfig.load) {
            try { await Layer8DConfig.load(); } catch (e) { console.error('Config load failed', e); }
        }

        var sidebarItems = document.querySelectorAll('.sidebar-item[data-section]');
        sidebarItems.forEach(function(item) {
            item.addEventListener('click', function(e) {
                e.preventDefault();
                var section = item.getAttribute('data-section');
                sidebarItems.forEach(function(s) { s.classList.remove('active'); });
                item.classList.add('active');
                if (typeof loadSection === 'function') loadSection(section);
            });
        });

        if (typeof loadSection === 'function') {
            var hashParts = typeof getHashParts === 'function' ? getHashParts() : { section: '', service: '' };
            var initSection = hashParts.section && sections[hashParts.section] ? hashParts.section : 'content';
            loadSection(initSection);
            if (hashParts.service) {
                setTimeout(function() {
                    var navItem = document.querySelector('.l8-subnav-item[data-service="' + hashParts.service + '"]');
                    if (navItem) navItem.click();
                }, 500);
            }
        }
    });
})();
