const {
    Record,
    StoreOf,
    Component,
    ListOf,
} = window.Torus;

const DATA_ORIGIN = '/data';

const PAGIATE_BY = 100;

const TODAY_ISO = (new Date()).toISOString().slice(0, 10);
