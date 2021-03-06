// @flow
import type {TypedState} from '../reducer'

export type LogTransformer = (action: TypedAction<*, *, *>) => Object

export type TypedAction<T, P, E> = {
  error?: false,
  type: T,
  payload: P,
  logTransformer?: LogTransformer,
} | {
  error: true,
  type: T,
  payload: E,
  logTransformer?: LogTransformer,
}

export type NoErrorTypedAction<T, P> = TypedAction<T, P, P>

export type Action = TypedAction<any, any, any>
export type GetState = () => TypedState
export type AsyncAction = (dispatch: Dispatch, getState: GetState) => ?Promise<*>
export type Dispatch = (action: AsyncAction | Action) => ?Promise<*>

export type TypedAsyncAction<A> = (dispatch: TypedDispatch<A>, getState: GetState) => ?Promise<*>
export type TypedDispatch<-A> = (action: TypedAsyncAction<A> | A) => ?Promise<*>

export const noPayloadTransformer: LogTransformer = (action) => {
  return {...action, payload: undefined}
}
