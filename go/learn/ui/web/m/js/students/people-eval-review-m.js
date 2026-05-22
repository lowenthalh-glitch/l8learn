/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Mobile finding review popup — thin wrapper around L8EvalReview core.
Include in m/app.html when mobile UI is implemented.
Requires: people-eval-review-core.js (shared), Layer8MPopup, Layer8MAuth.
*/
(function() {
    'use strict';

    window.L8EvalReviewMobile = {

        open: function(evalImport, serviceConfig) {
            var content = L8EvalReview.renderFindingsHtml(evalImport);

            Layer8MPopup.show({
                title: 'Review Evaluation Findings',
                content: content,
                size: 'full',
                showFooter: evalImport.processingStatus === 3,
                saveButtonText: 'Apply to Profile',
                onSave: function(popup) {
                    var payload = L8EvalReview.buildApplyPayload(evalImport);
                    var endpoint = Layer8MConfig.resolveEndpoint(serviceConfig.endpoint);
                    Layer8MAuth.put(endpoint, payload).then(function(resp) {
                        Layer8MUtils.showSuccess('Findings applied to profile');
                        Layer8MPopup.close();
                    }).catch(function() {
                        Layer8MUtils.showError('Failed to apply findings');
                    });
                },
                onShow: function(popup) {
                    popup.body.addEventListener('click', function(e) {
                        var btn = e.target.closest('[data-action]');
                        if (!btn) return;
                        var idx = parseInt(btn.getAttribute('data-idx'));
                        var action = btn.getAttribute('data-action');
                        L8EvalReview.onFindingAction(evalImport, idx, action);
                        popup.body.querySelector('.l8-eval-review').outerHTML = L8EvalReview.renderFindingsHtml(evalImport);
                    });
                }
            });
        }
    };
})();
