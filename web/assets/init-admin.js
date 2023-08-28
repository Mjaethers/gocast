document.addEventListener("alpine:init", () => {
    Alpine.directive("bind-change-set", (el, { expression, value, modifiers }, { evaluate, cleanup }) => {
        const changeSet = evaluate(expression);
        const fieldName = value;

        if (el.type === "file") {
            const [type] = modifiers;
            const isSingle = type === "single";

            const changeHandler = (e) => {
                changeSet.patch(fieldName, isSingle ? e.target.files[0] : e.target.files);
            };

            const onChangeSetUpdateHandler = (data) => {
                if (!data[fieldName]) {
                    el.value = "";
                }
            };

            changeSet.listen(onChangeSetUpdateHandler);
            el.addEventListener('change', changeHandler);

            cleanup(() => {
                changeSet.removeListener(onChangeSetUpdateHandler);
                el.removeEventListener('change', changeHandler)
            })
        } else {
            const keyupHandler = (e) => {
                changeSet.patch(fieldName, e.target.value);
            };

            const onChangeSetUpdateHandler = (data) => {
                el.value = data[fieldName];
            };

            changeSet.listen(onChangeSetUpdateHandler);
            el.addEventListener('keyup', keyupHandler)
            el.value = changeSet.get()[fieldName];

            cleanup(() => {
                changeSet.removeListener(onChangeSetUpdateHandler);
                el.removeEventListener('keyup', keyupHandler)
            })
        }
    });
});