/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Mobile batch-process wrapper. Uses L8BatchProcess core.
Include in m/app.html when mobile UI is implemented.
*/
(function() {
    'use strict';

    window.L8BatchProcessMobile = {

        openMultiUpload: function() {
            var refHtml = '<label>Student</label><input type="text" id="l8-batch-student" placeholder="Select student...">';
            var content = L8BatchProcess.renderUploadFormHtml(refHtml);

            Layer8MPopup.show({
                title: 'Upload Evaluations',
                content: content,
                size: 'full',
                showFooter: true,
                saveButtonText: 'Upload All',
                onSave: async function(popup) {
                    var studentInput = popup.body.querySelector('#l8-batch-student');
                    var studentId = studentInput ? studentInput.dataset.refId : '';
                    if (!studentId) {
                        Layer8MUtils.showError('Please select a student');
                        return;
                    }
                    var fileInput = popup.body.querySelector('#l8-batch-files');
                    if (!fileInput || !fileInput.files || fileInput.files.length === 0) {
                        Layer8MUtils.showError('Please select at least one PDF file');
                        return;
                    }

                    var endpoint = Layer8MConfig.resolveEndpoint('/20/EvalImprt');
                    var results = await L8BatchProcess.uploadFiles(
                        fileInput.files, studentId,
                        function(file) { return Layer8FileUpload.upload(file, 'general', '1'); },
                        function(evalData) { return Layer8MAuth.post(endpoint, evalData); },
                        null
                    );

                    Layer8MPopup.close();
                    if (results.success > 0) Layer8MUtils.showSuccess(results.success + ' file(s) uploaded');
                    if (results.failed > 0) Layer8MUtils.showError(results.failed + ' file(s) failed');
                }
            });
        },

        processAll: async function(tableData) {
            var endpoint = Layer8MConfig.resolveEndpoint('/20/EvalImprt');
            var result = await L8BatchProcess.triggerProcessAll(
                tableData,
                function(evalRecord) { return Layer8MAuth.put(endpoint, evalRecord); }
            );
            if (result.triggered) Layer8MUtils.showSuccess(result.message);
            else Layer8MUtils.showError(result.message);
        }
    };
})();
