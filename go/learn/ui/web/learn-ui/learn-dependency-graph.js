/*
 * L8Learn Module Dependency Graph
 * Overrides the default ERP dependency graph with l8learn-specific modules.
 */
(function() {
    'use strict';

    window.L8SysDependencyGraph = {
        modules: {
            content:    { label: 'Content',              icon: '📚', depends: [] },
            students:   { label: 'Students & People',    icon: '👨‍🎓', depends: [] },
            learning:   { label: 'Learning & Mastery',   icon: '🧠', depends: ['content', 'students'] },
            assessment: { label: 'Assessment',           icon: '📝', depends: ['content', 'students'] },
            analytics:  { label: 'Analytics & Reports',  icon: '📊', depends: ['students', 'learning', 'assessment'] },
            history:    { label: 'Historical Data',      icon: '📈', depends: ['students', 'analytics'] },
            collab:     { label: 'Collaboration',        icon: '🤝', depends: ['students'] },
            aimonitor:  { label: 'AI Monitor',           icon: '🤖', depends: [] }
        },

        subModules: {
            content: {
                'curriculum': { depends: [], foundation: true },
                'generated':  { depends: ['curriculum'] }
            },
            students: {
                'people':     { depends: [], foundation: true },
                'enrollment': { depends: ['people'] }
            },
            learning: {
                'adaptive': { depends: [], foundation: true }
            },
            assessment: {
                'results': { depends: [], foundation: true }
            },
            analytics: {
                'reports': { depends: [], foundation: true }
            },
            history: {
                'data': { depends: [], foundation: true }
            },
            collab: {
                'groups': { depends: [], foundation: true }
            },
            aimonitor: {
                'data': { depends: [], foundation: true }
            }
        },

        services: {}
    };

})();
