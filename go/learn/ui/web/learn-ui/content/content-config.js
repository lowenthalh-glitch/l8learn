(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Content',
        modules: {
            'curriculum': mod('Curriculum Framework', '📐', [
                svc('courses', 'Standards', '📕', '/10/Course', 'Course'),
                svc('units', 'Units', '📗', '/10/Unit', 'Unit'),
                svc('lessons', 'Lesson Plans', '📘', '/10/Lesson', 'Lesson'),
                svc('activities', 'Activity Bank', '🎯', '/10/Activity', 'Activity'),
                svc('worksheets', 'Worksheets', '📄', '/10/Worksheet', 'Worksheet')
            ]),
            'generated': mod('AI Generated', '🤖', [
                svc('genlessons', 'Generated Lessons', '🤖', '/10/GenLesson', 'GeneratedLesson')
            ])
        },
        submodules: ['ContentCurriculum', 'ContentGenerated']
    });
})();
