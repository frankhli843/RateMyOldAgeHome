import { createSlice, PayloadAction, createStore, combineReducers } from '@reduxjs/toolkit'

interface CounterState {
  value: number
}

const initialState = { value: 0 } as CounterState

const counter = createSlice({
  name: 'counter',
  initialState,
  reducers: {
    increment(state) {
      state.value++
    },
    decrement(state) {
      state.value--
    },
    incrementByAmount(state, action: PayloadAction<number>) {
      state.value += action.payload
    },
  },
})

const reducer = combineReducers({
    counter: counter.reducer,
  })
  
const store = createStore(reducer)

export const { 
    increment, 
    decrement, 
    incrementByAmount 
} = counter.actions
export default store;