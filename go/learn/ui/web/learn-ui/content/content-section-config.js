(function() {
    'use strict';
    Layer8SectionConfigs.register('content', {
        title: 'Content',
        subtitle: 'Curriculum & Learning Materials',
        icon: '📚',
        initFn: 'initializeContent',
        modules: [{
            key: 'curriculum',
            label: 'Curriculum',
            icon: '📚',
            isDefault: true,
            services: [
                { key: 'courses', label: 'Courses', icon: '📕', isDefault: true },
                { key: 'units', label: 'Units', icon: '📗' },
                { key: 'lessons', label: 'Lessons', icon: '📘' },
                { key: 'activities', label: 'Activities', icon: '🎯' },
                { key: 'worksheets', label: 'Worksheets', icon: '📄' }
            ]
        }]
    });
})();
