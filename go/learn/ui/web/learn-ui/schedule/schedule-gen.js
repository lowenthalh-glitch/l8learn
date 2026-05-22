/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Desktop schedule generation wrapper.
*/
(function() {
    'use strict';

    function getHeaders() {
        return Object.assign({ 'Content-Type': 'application/json' },
            typeof getAuthHeaders === 'function' ? getAuthHeaders() : {});
    }

    window.L8ScheduleGenDesktop = {

        open: function(studentId) {
            var regConf = (typeof Layer8DReferenceRegistry !== 'undefined' && Layer8DReferenceRegistry.Student) || {};
            var refCfg = JSON.stringify({
                modelName: 'Student', endpoint: Layer8DConfig.resolveEndpoint('/20/Student'),
                idColumn: regConf.idColumn || 'studentId', displayColumn: regConf.displayColumn || 'lastName',
                selectColumns: regConf.selectColumns, displayLabel: 'Student'
            });
            var refHtml = '<label>Student</label>' +
                "<input type=\"text\" id=\"l8-sched-student\" name=\"studentId\" class=\"reference-input\" " +
                "data-ref-config='" + Layer8DUtils.escapeAttr(refCfg) + "' " +
                "data-ref-id=\"" + (studentId || '') + "\" data-lookup-model=\"Student\" " +
                "data-field-key=\"studentId\" readonly placeholder=\"Click to select...\">";

            var content = L8ScheduleGen.renderFormHtml(refHtml);

            Layer8DPopup.show({
                title: 'Generate Weekly Schedule',
                content: content,
                size: 'medium',
                showFooter: true,
                saveButtonText: 'Generate',
                onSave: async function() {
                    var body = Layer8DPopup.getBody();
                    var input = body.querySelector('#l8-sched-student');
                    var sid = input ? input.dataset.refId : studentId;
                    if (!sid) {
                        Layer8DNotification.error('Please select a student');
                        return;
                    }
                    var data = L8ScheduleGen.collectFormData(body, sid);
                    try {
                        var resp = await fetch(Layer8DConfig.resolveEndpoint('/30/Schedule'), {
                            method: 'POST', headers: getHeaders(), body: JSON.stringify(data)
                        });
                        if (!resp.ok) throw new Error('POST failed');
                        Layer8DPopup.close();
                        Layer8DNotification.success('Schedule generation started. This may take several minutes...');
                        // Open schedule view after a delay to let blocks generate
                        var result = await resp.json();
                        if (result && result.id) {
                            setTimeout(function() { L8ScheduleViewDesktop.openByQuery(sid); }, 5000);
                        }
                    } catch (err) {
                        Layer8DNotification.error('Failed to generate schedule: ' + err.message);
                    }
                },
                onShow: function(body) {
                    setTimeout(function() { Layer8DForms.attachReferencePickers(body); }, 50);
                }
            });
        }
    };
})();
