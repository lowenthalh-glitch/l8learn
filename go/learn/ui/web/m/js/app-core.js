/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/**
 * Mobile L8Learn - App Core - Navigation and initialization
 */
(function() {
    'use strict';

    const SECTIONS = {
        'dashboard': 'sections/dashboard.html',
        'system': 'sections/system.html'
    };

    let currentSection = 'dashboard';
    let sectionCache = {};

    window.showErrorAndLogout = function(message, detail) {
        if (typeof Layer8MAuth !== 'undefined') {
            Layer8MAuth.showErrorAndLogout(message, detail);
        } else {
            alert(message + (detail ? '\n\n' + detail : ''));
            window.location.href = '../login.html';
        }
    };

    window.MobileApp = {
        async init() {
            if (!Layer8MAuth.requireAuth()) return;

            await Layer8MConfig.load();
            await Layer8DConfig.load();

            this.updateUserInfo();

            var token = Layer8MAuth.getBearerToken();

            try {
                var permResp = await fetch('/permissions', {
                    headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' }
                });
                if (permResp.ok) {
                    window.Layer8DPermissions = await permResp.json();
                }
            } catch (e) { console.warn('Failed to load permissions:', e); }

            if (typeof Layer8DModuleFilter !== 'undefined') {
                var configLoaded = await Layer8DModuleFilter.load(token);
                if (!configLoaded) return;
                this.applyModuleFilter();
            }

            if (typeof Layer8MPortalSwitcher !== 'undefined') {
                var headerActions = document.querySelector('.header-actions');
                if (headerActions) {
                    Layer8MPortalSwitcher.init({
                        container: headerActions,
                        insertBefore: document.getElementById('refresh-btn'),
                        apiPrefix: Layer8MConfig.getApiPrefix(),
                        currentPath: window.location.pathname
                    });
                }
            }

            this.initSidebar();

            document.getElementById('refresh-btn')?.addEventListener('click', function() {
                MobileApp.loadSection(currentSection, true);
            });

            var hash = window.location.hash.slice(1);
            var section = SECTIONS[hash] ? hash : 'dashboard';
            await this.loadSection(section);

            window.addEventListener('hashchange', function() {
                var newSection = window.location.hash.slice(1);
                if (SECTIONS[newSection] && newSection !== currentSection) {
                    MobileApp.loadSection(newSection);
                }
            });
        },

        updateUserInfo() {
            var username = Layer8MAuth.getUsername();
            var initial = username.charAt(0).toUpperCase();
            document.getElementById('user-name').textContent = username;
            document.getElementById('user-avatar').textContent = initial;
        },

        initSidebar() {
            var menuToggle = document.getElementById('menu-toggle');
            var overlay = document.getElementById('sidebar-overlay');

            if (menuToggle) menuToggle.addEventListener('click', function() { MobileApp.openSidebar(); });
            if (overlay) overlay.addEventListener('click', function() { MobileApp.closeSidebar(); });

            document.querySelectorAll('.sidebar-item[data-section]').forEach(function(item) {
                item.addEventListener('click', async function(e) {
                    e.preventDefault();
                    var section = item.dataset.section;
                    var module = item.dataset.module;
                    MobileApp.closeSidebar();
                    await MobileApp.loadSection(section);
                    if (module && window.Layer8MNav) {
                        Layer8MNav.navigateToModule(module);
                    }
                });
            });
        },

        openSidebar() {
            var sidebar = document.getElementById('sidebar');
            var overlay = document.getElementById('sidebar-overlay');
            if (sidebar) sidebar.classList.add('open');
            if (overlay) overlay.classList.add('visible');
            document.body.style.overflow = 'hidden';
        },

        closeSidebar() {
            var sidebar = document.getElementById('sidebar');
            var overlay = document.getElementById('sidebar-overlay');
            if (sidebar) sidebar.classList.remove('open');
            if (overlay) overlay.classList.remove('visible');
            document.body.style.overflow = '';
        },

        async loadSection(section, forceReload) {
            if (section !== 'dashboard' && window.LAYER8M_NAV_CONFIG && LAYER8M_NAV_CONFIG[section]) {
                await this._loadDashboardForModule(section, forceReload);
                return;
            }

            var sectionUrl = SECTIONS[section];
            if (!sectionUrl) {
                console.error('Unknown section:', section);
                return;
            }

            this.updateNavState(section);

            var contentArea = document.getElementById('content-area');
            if (!contentArea) return;

            contentArea.style.opacity = '0.5';

            try {
                if (!forceReload && sectionCache[section]) {
                    contentArea.innerHTML = sectionCache[section];
                } else {
                    var response = await fetch(sectionUrl + '?t=' + Date.now());
                    if (!response.ok) throw new Error('Failed to load section');
                    var html = await response.text();
                    sectionCache[section] = html;
                    contentArea.innerHTML = html;
                }

                this.executeScripts(contentArea);
                this.initSection(section);

                currentSection = section;
                window.location.hash = section;
                contentArea.scrollTop = 0;
            } catch (error) {
                console.error('Error loading section:', error);
                contentArea.innerHTML =
                    '<div class="empty-state">' +
                    '<span class="empty-state-icon">&#x26A0;</span>' +
                    '<h4 class="empty-state-title">Failed to load</h4>' +
                    '<p class="empty-state-message">Please try again</p>' +
                    '<button class="btn btn-primary" onclick="MobileApp.loadSection(\'' + section + '\', true)">Retry</button>' +
                    '</div>';
            }

            contentArea.style.opacity = '1';
        },

        async _loadDashboardForModule(moduleKey, forceReload) {
            this.updateNavState(moduleKey);

            var contentArea = document.getElementById('content-area');
            if (!contentArea) return;

            contentArea.style.opacity = '0.5';

            try {
                if (!forceReload && sectionCache['dashboard']) {
                    contentArea.innerHTML = sectionCache['dashboard'];
                } else {
                    var response = await fetch(SECTIONS['dashboard'] + '?t=' + Date.now());
                    if (!response.ok) throw new Error('Failed to load dashboard');
                    var html = await response.text();
                    sectionCache['dashboard'] = html;
                    contentArea.innerHTML = html;
                }

                this.executeScripts(contentArea);
                this.initSection('dashboard');

                Layer8MNav.navigateToModule(moduleKey);

                currentSection = moduleKey;
                window.location.hash = moduleKey;
                contentArea.scrollTop = 0;
            } catch (error) {
                console.error('Error loading module:', error);
            }

            contentArea.style.opacity = '1';
        },

        updateNavState(section) {
            document.querySelectorAll('.sidebar-item').forEach(function(item) {
                item.classList.remove('active');
                if (item.dataset.section === section) {
                    item.classList.add('active');
                }
            });
        },

        executeScripts(container) {
            var scripts = container.querySelectorAll('script');
            scripts.forEach(function(oldScript) {
                var newScript = document.createElement('script');
                Array.from(oldScript.attributes).forEach(function(attr) {
                    newScript.setAttribute(attr.name, attr.value);
                });
                newScript.textContent = oldScript.textContent;
                oldScript.parentNode.replaceChild(newScript, oldScript);
            });
        },

        initSection(section) {
            var initFunctions = {
                'dashboard': 'initMobileDashboard',
                'system': 'initMobileSystem'
            };

            var initFn = initFunctions[section];
            if (initFn && typeof window[initFn] === 'function') {
                window[initFn]();
            }
        },

        getCurrentSection() {
            return currentSection;
        },

        logout() {
            Layer8MAuth.logout();
        },

        applyModuleFilter() {
            if (!window.Layer8DModuleFilter) return;
            document.querySelectorAll('.sidebar-item[data-section]').forEach(function(item) {
                var section = item.dataset.section;
                if (section === 'dashboard' || item.dataset.module === 'system') return;
                var moduleKey = item.dataset.module || section;
                if (!Layer8DModuleFilter.isEnabled(moduleKey)) {
                    item.style.display = 'none';
                }
            });
        },

        switchToDesktop() {
            localStorage.setItem('preferDesktop', 'true');
            window.location.href = '../';
        }
    };

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', function() { MobileApp.init(); });
    } else {
        MobileApp.init();
    }

})();
