package com.medclever.pocketdoctor_free.repositories

import android.content.res.AssetManager
import com.medclever.pocketdoctor_free.entities.Message
import com.medclever.pocketdoctor_free.entities.MessageData
import com.medclever.pocketdoctor_free.utils.readFileFromAssets
import kotlinx.serialization.json.*

class MessagesRepository(val assets: AssetManager, val root: String) {
    val json = Json {
        namingStrategy = JsonNamingStrategy.SnakeCase
    }

    public fun getAll(): List<Message> {
        val messagesText = readFileFromAssets(this.root + "/messages.json", this.assets)
        val listData: List<MessageData> = this.json.decodeFromString(messagesText)
        return listData.map{ m -> Message(m.message) }
    }
}