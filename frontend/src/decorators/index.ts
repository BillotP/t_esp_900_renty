import { createDecorator, VueDecorator } from 'vue-class-component';
import { WatchOptions, PropOptions } from 'vue';

export type Constructor = new(...args: any[]) => any;

export function Props(options: (PropOptions | Constructor[] | Constructor) = {}): VueDecorator {
    return createDecorator((componentOptions, k) => {
        (componentOptions.props || (componentOptions.props = {}) as any)[k] = options;
    });
}

export const Prop: VueDecorator = createDecorator((componentOptions, k) => {
    (componentOptions.props || (componentOptions.props = {}) as any)[k] = {};
});

export function Watch(path: string, options: WatchOptions = {}): VueDecorator {
    const { deep = false, immediate = false } = options;

    return createDecorator((componentOptions, handler) => {
        if (typeof componentOptions.watch !== 'object') {
            componentOptions.watch = Object.create(null);
        }

        const watch: any = componentOptions.watch;

        if (typeof watch[path] === 'object' && !Array.isArray(watch[path])) {
            watch[path] = [watch[path]];
        } else if (typeof watch[path] === 'undefined') {
            watch[path] = [];
        }

        watch[path].push({ handler, deep, immediate });
    });
}
