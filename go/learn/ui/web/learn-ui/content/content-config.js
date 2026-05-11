(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Content',
        modules: {
            'curriculum': mod('Curriculum', '📚', [
                svc('courses', 'Courses', '📕', '/10/Course', 'Course'),
                svc('units', 'Units', '📗', '/10/Unit', 'Unit'),
                svc('lessons', 'Lessons', '📘', '/10/Lesson', 'Lesson'),
                svc('activities', 'Activities', '🎯', '/10/Activity', 'Activity'),
                svc('worksheets', 'Worksheets', '📄', '/10/Worksheet', 'Worksheet')
            ])
        },
        submodules: ['ContentCurriculum']
    });
})();
