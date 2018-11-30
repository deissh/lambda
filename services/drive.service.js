"use strict";
const path = require("path");
const mkdir = require("mkdirp").sync;
const StorageManager = require('@slynova/flydrive');

module.exports = {
  name: 'lambda-drive',
  storage: null,
  setting: {
    STORAGE_ROOT: './upload',
    storageConfig: null
  },
  actions: {
    /**
     *
     * Get a disk storage
     *
     * @param {String} name - the disk name
     *
     * @Returns {Object} - A storage instance with the associated driver
     */
    disk(name) {
      return name ? this.storage.disk(name) : this.storage.disk();
    },
    /**
     * Add new Driver
     *
     * @param {String} name - The driver name
     * @param {Object|function} driver - driver contructor function
     * @param {Object?} config - Thre configuration to use for the driver construction function
     *
     * @Returns {Void}
     */
    extends(name, driver) {
      this.storage.extend(name, driver);
    }
  },
  async created() {
    // Create data folder
    this.settings.STORAGE_ROOT = this.settings.STORAGE_ROOT || "";
    this.settings.STORAGE_ROOT = (this.settings.STORAGE_ROOT.replace(/(\/$)|(\\$)/, ''));
    if (!this.settings.storageConfig) {
      this.settings.storageConfig = {
        default: 'local',
        disks: {
          local: {
            driver: 'local',
            root: this.settings.STORAGE_ROOT
          }
        }
      }
    }
    this.settings.storageConfig.disks =  Object.assign({
      local: {
        driver: 'local',
        root: this.settings.STORAGE_ROOT
      }
    },this.settings.storageConfig.disks);

    this.settings.storageConfig.disks.local = Object.assign({
      driver: 'local',
      root: this.settings.STORAGE_ROOT
    },this.settings.storageConfig.disks.local);
    //if a root was already defined in config
    this.settings.STORAGE_ROOT = (this.settings.storageConfig.disks.local.root);
    this.settings.storageConfig.default = this.settings.defaultStorage || 'local';
    this.settings.defaultStorage = this.settings.storageConfig.default;
    if (tryToCreateRoot) {
      mkdir(this.settings.STORAGE_ROOT);
    }

    this.storage = new StorageManager(this.settings.storageConfig);
    //register custom drivers definition
    this.settings.storageConfig.customDrivers = Object.assign({},this.settings.storageConfig.customDrivers);
    for(const name in this.settings.storageConfig.customDrivers){
      this.extends(name,this.settings.storageConfig.customDrivers[name]);
      console.log("name", name);
    }
  }
}