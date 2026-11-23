# manager-life

Một API hello world bằng Go, dùng để học CI/CD với Jenkins.

## Chạy

```bash
make run     # chạy bằng Go
make test    # chạy test
make up      # chạy bằng Docker
```

Thử: `curl localhost:8081/hello`

## Jenkins

Pipeline trong [Jenkinsfile](Jenkinsfile): **Test -> Build -> Deploy -> Check**.

1. Tạo git repo (Jenkins chỉ thấy code đã commit):

   ```bash
   git init -b main
   git add .
   git commit -m "Initial commit"
   ```

2. Chạy Jenkins và lấy mật khẩu mở khoá:

   ```bash
   make jenkins-up
   make jenkins-password
   ```

3. Mở <http://localhost:8080>, dán mật khẩu, chọn **Install suggested plugins**, tạo tài khoản admin.

4. **New Item** -> tên `manager-life` -> **Pipeline**. Ở mục Pipeline điền:

   | Ô | Giá trị |
   | --- | --- |
   | Definition | Pipeline script from SCM |
   | SCM | Git |
   | Repository URL | `file:///repo` |
   | Branch Specifier | `*/main` |
   | Script Path | `Jenkinsfile` |

5. **Save** -> **Build Now** -> bấm vào số build -> **Console Output**.

Từ lần sau: sửa code, `git commit`, khoảng 2 phút sau Jenkins tự build và deploy.
