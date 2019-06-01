# go-fmt-func-sig

Formats function signatures into newlines when there are many parameters.

## Example

```diff
$ diff -u foo/bar.go <(go run github.com/execjosh/go-fmt-func-sig < foo/bar.go)
--- foo/bar.go	2019-06-01 23:55:48.000000000 +0900
+++ /dev/fd/63	2019-06-02 00:14:54.000000000 +0900
@@ -10,11 +10,17 @@
 	ctx context.Context,
 	name string,
 	message string,
-	count int) (int, error) {
+	count int,
+) (int, error) {
 	return 0, nil
 }

-func notFormatted(fset *token.FileSet, fl *ast.FieldList, f *ast.File, a, b, c string) {
+func notFormatted(
+	fset *token.FileSet,
+	fl *ast.FieldList,
+	f *ast.File,
+	a, b, c string,
+) {
 }

 func alreadyFormatted(
@@ -25,4 +31,10 @@
 ) {
 }

-func finalLineInFile(fset *token.FileSet, fl *ast.FieldList, f *ast.File, a, b, c string) {}
+func finalLineInFile(
+	fset *token.FileSet,
+	fl *ast.FieldList,
+	f *ast.File,
+	a, b, c string,
+) {
+}
```
