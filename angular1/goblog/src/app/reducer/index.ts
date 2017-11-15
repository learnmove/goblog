import { counterReducer } from "./counter";
import { userReducer } from "./user";
export const rootReducer = {
        counter:counterReducer,
        user:userReducer,
      }