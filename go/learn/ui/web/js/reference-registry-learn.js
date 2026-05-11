/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 * Licensed under the Apache License, Version 2.0
 */
(function() {
    'use strict';
    var ref = window.Layer8RefFactory;
    Layer8DReferenceRegistry.register({
        ...ref.simple('Course', 'courseId', 'name', 'Course'),
        ...ref.simple('Unit', 'unitId', 'name', 'Unit'),
        ...ref.simple('Lesson', 'lessonId', 'name', 'Lesson'),
        ...ref.simple('Activity', 'activityId', 'name', 'Activity'),
        ...ref.simple('Worksheet', 'worksheetId', 'name', 'Worksheet'),
        ...ref.simple('Skill', 'skillId', 'name', 'Skill'),
        ...ref.simple('Student', 'studentId', 'lastName', 'Student'),
        ...ref.simple('Guardian', 'guardianId', 'lastName', 'Guardian'),
        ...ref.simple('Teacher', 'teacherId', 'lastName', 'Teacher'),
        ...ref.simple('Classroom', 'classroomId', 'name', 'Classroom'),
        ...ref.simple('School', 'schoolId', 'name', 'School'),
        ...ref.simple('District', 'districtId', 'name', 'District'),
        ...ref.simple('Family', 'familyId', 'name', 'Family'),
        ...ref.simple('LearningPod', 'podId', 'name', 'Pod'),
        ...ref.simple('LearningPath', 'pathId', 'pathId', 'Path'),
        ...ref.simple('Benchmark', 'benchmarkId', 'name', 'Benchmark'),
        ...ref.simple('CollabGroup', 'groupId', 'name', 'Group'),
        ...ref.simple('Challenge', 'challengeId', 'name', 'Challenge'),
        ...ref.simple('StudentProfile', 'profileId', 'profileId', 'Profile'),
        ...ref.simple('EvalImport', 'importId', 'professionalName', 'Evaluation'),
        ...ref.simple('LLMPromptLog', 'logId', 'logId', 'Prompt Log'),
        ...ref.simple('LLMConfig', 'configId', 'configId', 'LLM Config'),
        ...ref.simple('GeneratedLesson', 'generatedLessonId', 'title', 'Generated Lesson')
    });
})();
