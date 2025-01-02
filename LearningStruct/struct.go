package main

import "fmt"

type Student struct { 
	FirstName string
	LastName  string
	Std       int
	Roll      int
}

func main() {
	// 1st way of assign value to Studnet dataTypes
	var student Student // declaring student variable of Student datatypes
	student.FirstName = "Himanshu"
	student.LastName = "Raj"
	student.Std = 12
	student.Roll = 1

	fmt.Println("student 1st name is :", student.FirstName) // printing single attribute of student
	// print complete details of student
	fmt.Println("Details of student is :", student)

	//Direct using Student datatype faster way
	student1 := Student{
		FirstName: "Viany",
		LastName:  "Khardoniya",
		Std:       12,
		Roll:      2,
	}
	fmt.Println("Details of student1 is :", student1)

	//assign value using new keyword
	var student2 = new(Student)
	student2.FirstName = "Anurag"
	student2.LastName = "patel"
	student2.Std = 12
	// student2.Roll =3 // if we not assign Roll then defualt value of its datatype(0)will assign
	fmt.Println("Detials of student2 is :", student2) //The & symbol before the struct(output) indicates that student2 is a pointer to the struct.

	// nested struct
	type details struct {
		mobile  string
		address string
		state   string
	}

	type student1_info struct {
		basic_Details Student
		extra_Details details // extra_Details variable will store other struct detais type vlaue
	}

	//assing vlaue for student_info
	var student_details student1_info

	student_details.basic_Details = Student{ // 1st attribute of student_info
		FirstName: "Nilesh",
		LastName:  "Singh",
		Std:       11,
		Roll:      5,
	}
	// student_details.extra_Details.address = "Motihari" // acccess single feild
	student_details.extra_Details = details{
		mobile:  "123456",
		address: "Motihari",
		state:   "Bihar",
	}

	fmt.Println("Complete details for Student_information is :",student_details) // printing complete details 
	student_details.extra_Details.mobile ="9874652453" // changing feild
	fmt.Println("After change mobile details is :" , student_details.extra_Details.mobile) // printing single feilds

}
