(function() {
    'use strict';
    Layer8SectionConfigs.register('content', {
        title: 'Content',
        subtitle: 'Curriculum Standards & AI-Generated Lessons',
        icon: '📚',
        initFn: 'initializeContent',
        modules: [
            {
                key: 'curriculum', label: 'Curriculum Framework', icon: '📐', isDefault: true,
                services: [
                    { key: 'courses', label: 'Standards', icon: '📕', isDefault: true },
                    { key: 'units', label: 'Units', icon: '📗' },
                    { key: 'lessons', label: 'Lesson Plans', icon: '📘' },
                    { key: 'activities', label: 'Activity Bank', icon: '🎯' },
                    { key: 'worksheets', label: 'Worksheets', icon: '📄' }
                ]
            },
            {
                key: 'generated', label: 'AI Generated', icon: '🤖',
                services: [
                    { key: 'genlessons', label: 'Generated Lessons', icon: '🤖', isDefault: true }
                ]
            }
        ]
    });
})();
