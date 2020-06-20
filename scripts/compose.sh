#!/bin/zsh

DOCKER_PATH=./build/docker-compose.yaml
PROJECT=cluster-executor

DM="docker-compose -f $DOCKER_PATH -p $PROJECT"

function up() {
    echo $DM ps
    sh -c "${DM} ps"

    echo $DM pull
    sh -c "${DM} pull"

    echo $DM up writer
    sh -c "${DM} up -d db"
    sh -c "${DM} run --rm wait"
    if [ $? -gt 0 ]; then
        exit 1
    fi

    echo $DM up reader
    sh -c "${DM} up -d db-reader"
    sh -c "${DM} run --rm wait-reader"
    if [ $? -gt 0 ]; then
        exit 1
    fi

    echo $DM ps
    sh -c "${DM} ps"
}

function down() {
    echo $DM down
    sh -c "${DM} down"
}

function clean() {
    echo $DM clean
    sh -c "${DM} down --volumes"
}

function log() {
    container=$1

    echo $DM down
    sh -c "${DM} logs $1"
}

function ps() {
    echo $DM down
    sh -c "${DM} ps"
}

function exec() {
    echo $DM exec
    sh -c "${DM} exec db mysql -uroot -p"
}

$1 $2
