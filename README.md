{
scores(id: 1){
id
score
calculateDate
scoreGrade
    }
}


{
company(id: 1){
id
name
address
scores{
id
score
calculateDate
scoreGrade
}
}
}