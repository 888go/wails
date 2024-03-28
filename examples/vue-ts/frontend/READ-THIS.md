# 该模板使用了一个变通方法，因为默认模板由于以下问题无法编译：

https://github.com/vuejs/core/issues/1228

在`tsconfig.json`文件中，将`isolatedModules`设置为`false`而非`true`，以此作为解决该问题的临时方案。
