export const UserHelper={
    getUser:()=>{
      return  JSON.parse(localStorage.getItem("user")).user
    },
    setUser:(data)=>{
      localStorage.setItem("user",JSON.stringify(data))  
    }
}