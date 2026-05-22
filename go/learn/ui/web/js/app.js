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

        // Load per-type action permissions for the current user
        try {
            var permResp = await fetch('/permissions', {
                headers: { 'Authorization': 'Bearer ' + bearerToken, 'Content-Type': 'application/json' }
            });
            if (permResp.ok) {
                window.Layer8DPermissions = await permResp.json();
            }
        } catch (e) { console.warn('Failed to load permissions:', e); }

        // Apply permission-based nav filtering
        if (typeof Layer8DPermissionFilter !== 'undefined') {
            var nsMap = {
                'content': 'Content',
                'students': 'Students',
                'learning': 'Learning',
                'assessment': 'Assessment',
                'analytics': 'Analytics',
                'history': 'History',
                'collab': 'Collab',
                'aimonitor': 'AIMonitor'
            };

            Layer8DPermissionFilter.registerResolver(function(sectionKey, moduleKey, serviceKey) {
                var nsName = nsMap[sectionKey];
                var ns = nsName ? window[nsName] : null;
                if (!ns || !ns.modules || !ns.modules[moduleKey]) return null;
                var svc = ns.modules[moduleKey].services.find(function(s) { return s.key === serviceKey; });
                return svc ? svc.model : null;
            });

            var sidebarModels = {};
            Object.keys(nsMap).forEach(function(section) {
                var ns = window[nsMap[section]];
                if (!ns || !ns.modules) return;
                var models = [];
                Object.values(ns.modules).forEach(function(mod) {
                    if (mod.services) mod.services.forEach(function(svc) { if (svc.model) models.push(svc.model); });
                });
                sidebarModels[section] = models;
            });
            Layer8DPermissionFilter.applyToSidebar(sidebarModels);
        }

        var sidebarItems = document.querySelectorAll('.sidebar-item[data-section]');
        sidebarItems.forEach(function(item) {
            item.addEventListener('click', function(e) {
                e.preventDefault();
                var section = item.getAttribute('data-section');
                sidebarItems.forEach(function(s) { s.classList.remove('active'); });
                item.classList.add('active');
                if (typeof loadSection === 'function') {
                    loadSection(section);
                    setTimeout(function() {
                        if (typeof Layer8DPermissionFilter !== 'undefined') {
                            Layer8DPermissionFilter.applyToSection(section);
                        }
                    }, 300);
                }
            });
        });

        if (typeof loadSection === 'function') {
            var hashParts = typeof getHashParts === 'function' ? getHashParts() : { section: '', service: '' };
            var initSection = hashParts.section && sections[hashParts.section] ? hashParts.section : 'students';
            loadSection(initSection);
            setTimeout(function() {
                if (typeof Layer8DPermissionFilter !== 'undefined') {
                    Layer8DPermissionFilter.applyToSection(initSection);
                }
            }, 300);
            if (hashParts.service) {
                setTimeout(function() {
                    var navItem = document.querySelector('.l8-subnav-item[data-service="' + hashParts.service + '"]');
                    if (navItem) navItem.click();
                }, 500);
            }
        }
    });
})();
