DROP DATABASE IF EXISTS class_roster;
CREATE DATABASE class_roster;

USE class_roster;

CREATE TABLE student (
	studentID int primary key auto_increment,
    firstName varchar(48) NOT NULL,
    lastName varchar(48) NOT NULL
);

CREATE TABLE teacher (
	teacherID int primary key auto_increment,
    firstName varchar(48) NOT NULL,
    lastName varchar(48) NOT NULL,
    title varchar(5),
    office varchar(10),
    department varchar(48)
);

CREATE TABLE course (
	courseID int primary key auto_increment,
    courseName varchar(48) NOT NUll
);

CREATE TABLE studentcourse (
	studentID int,
    courseID int,
    semester varchar(25),
    grade varchar(1),
    primary key(studentID, courseID),
    foreign key fk_studentID (studentID) references student(studentID) ON DELETE CASCADE,
    foreign key fk_courseID (courseID) references course(courseID) ON DELETE CASCADE
);

CREATE TABLE teachercourse (
	teacherID int,
    courseID int,
    semester varchar(25),
    primary key(teacherID, courseID),
    foreign key fk_teacherID (teacherID) references teacher(teacherID) ON DELETE CASCADE,
    foreign key fk_courseID (courseID) references course(courseID) ON DELETE CASCADE
);

-- Data imports for the normalization example
INSERT INTO course(courseName) VALUES 
("Math"),
("History"),
("Java"),
("Python"),
("Python"),
("C#");

INSERT INTO student(firstName, lastName) VALUES
("Alice", "Smith"),
("Bob", "Jones"),
("Carol", "White"),
("Emma", "Thompson"),
("Frank", "Rivera"),
("Grace", "Kim"),
("Henry", "Lopez");

INSERT INTO teacher(firstName, lastName, title, office, department) VALUES 
("Steven", "Even", "Dr", "Rm101", "Math"),
("Chun", "Lee", "Prof", "Rm202", "History"),
("Maman", "Aman", "Dr", "Rm307", "CompSci"),
("Ten", "Tan", "Dr", "Rm421", "Physics"),
("So", "Green", "Dr", "Rm222", "CompSci");

INSERT INTO studentcourse(studentID, courseID, semester, grade) VALUES
(1, "1", "Fall 2025", "A"), 
(1, "2", "Fall 2025", "B"),
(2, "1", "Fall 2025", "B"),
(3, "2", "Fall 2025", "A"),
(5, "3", "Fall 2025", "B"),
(6, "4", "Fall 2025", "A"), 
(7, "5", "Fall 2025", "A"),
(4, "6", "Fall 2025", "B"); 

INSERT INTO teachercourse(teacherID, courseID, semester) VALUES
(1, "1", "Fall 2025"),
(2, "2", "Fall 2025"),
(3, "3", "Fall 2025"), 
(4, "4", "Fall 2025"),
(4, "5", "Fall 2025"), 
(5, "6", "Fall 2025");

