# generateStuff

Application that generates some stuffs like birthdate, names, jobs, cities etc.

## Methods:

### GenerateBirthdate
internal function that generates birthdate thinking in the last 110 years. 
The function takes leap years, months with 30 days and 31 days into account.

### GenerateFirstName
internal function that generates first name randomly

### GenerateFirstNameWithProbabilities
internal function that generates first name. Names with more probabilities will have more chances to be generated.

### GenerateSurname
internal function that generates a surname randomly

### GenerateSurnameNameWithProbabilities
internal function that generates a surname. Surnames with more probabilities will have more chances to be generated.

### GenerateFullName
internal function that generates a full name. This function generates a first name and 2 surnames. Ex: Anthony Lewis Wisbech

### GenerateFullNameWirhProbabilities
internal function that generates a full name. This function generates a first name and 2 surnames using the functions GenerateFirstNameWithProbabilities and GenerateSurnameNameWithProbabilities