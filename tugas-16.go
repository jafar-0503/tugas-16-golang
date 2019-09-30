package main

import "fmt"
import "database/sql"
import _"mysql-master"

type mahasiswa struct{
  id int
  nama string
  jurusan string
  alamat string
}

func koneksi() (*sql.DB, error){
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/materi_golang")
  if err != nil{
    return nil, err
  }
    return db, nil
}

func sql_tampil(){
  db, err := koneksi()
    if err != nil{
      fmt.Println(err.Error())
      return
    }
    defer db.Close()
    rows, err := db.Query("select * from tbl_mahasiswa")
    if err != nil{
      fmt.Println(err.Error())
      return
    }
    defer rows.Close()

    var result []mahasiswa

    for rows.Next(){
      var each = mahasiswa{}

      var err = rows.Scan(&each.id, &each.nama, &each.jurusan, &each.alamat)

      if err != nil{
        fmt.Println(err.Error())
        return
      }
      result = append(result, each)
    }

    if err = rows.Err(); err != nil{
      fmt.Println(err.Error())
      return
    }
    for _, each:= range result{
      fmt.Println(each.nama, each.jurusan, each.alamat)
    }
}

func main(){
  sql_tampil()
}
