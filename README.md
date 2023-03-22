# Go-Cache-API

## 1. Pull the docker image :
```
  docker pull siddheshk02/api-dev
```

## 2. Run the image :

```
  docker run -p 8000:8000 api-dev
```
## Output :
![image](https://user-images.githubusercontent.com/90148705/226841504-9994a2b5-eead-4f6a-9393-a9486153bfbd.png)

3. Now head over to ``` http://127.0.0.1:8000/ ``` on your browser, You'll see ``` Hello, World 👋! ```.

4. ``` /posts/:id ``` : Go to url ``` http://127.0.0.1:8000//posts/1 ``` 
   Note : you can try with any ``` id ```.
   
   Output : 
   
   ![image](https://user-images.githubusercontent.com/90148705/226843122-f2af5b41-6e57-4e75-9b02-d1953f80047b.png)
   
   You'll get some data like this.

5. ``` /todos/:id ``` : Go to url ``` http://127.0.0.1:8000//todos/1 ``` 
   Note : you can try with any ``` id ```.
   
   Output :   
   
   ![image](https://user-images.githubusercontent.com/90148705/226843854-01a53ce0-1845-420b-b6fd-1fcf78c83c17.png)
    
   You'll get some data like this.
